package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	utils "fileupload/utils"
)

// 给变量赋值
var (
	env           string
	maxUploadSize int64
	help          = pflag.BoolP("help", "h", false, "查看命令帮助")
)

func init() {
	pflag.StringVarP(&env, "env", "e", "test", "启动环境,可取值:dev、test、production")
	pflag.Int64VarP(&maxUploadSize, "maxSzie", "m", 3, "上传文件大小限制,单位为MB")
	pflag.Parse()
	maxUploadSize = maxUploadSize * 1024 * 1024
	// log.Println("pflag", env, maxUploadSize)
}

type TemplateParam struct {
	Host string
}

var (
	uploadPath    string
	port          string
	templateParam TemplateParam
)

func init() {
	viper.AddConfigPath("./config")      // 把当前目录加入到配置文件的搜索路径中
	viper.SetConfigName("config_" + env) // 配置文件名称（没有文件扩展名）
	viper.SetConfigType("yaml")          // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件

	if err := viper.ReadInConfig(); err != nil { // 读取配置文件。如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	uploadPath = viper.Get("uploadPath").(string)
	port = viper.Get("port").(string)
	templateParam.Host = viper.Get("host").(string)
	// log.Println("viper", uploadPath, templateParam.Host, port)
}

type HTTPHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func handler(h http.HandlerFunc, decorators ...HTTPHandlerDecorator) http.HandlerFunc {
	for i := range decorators {
		d := decorators[len(decorators)-i-1] // in reverse
		h = d(h)
	}
	return h
}

func main() {
	if *help {
		pflag.Usage()
		return
	}

	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/upload", handler(uploadFileHandler, preChecker))
	serverMux.Handle("/", shareServer(http.FileServer(http.Dir(uploadPath))))

	log.Printf("Server started on port %s, use /upload for uploading files", port)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: serverMux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("server start error: ", err)
	}
}

func preChecker(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 请求类型只能为POST
		if r.Method != "POST" {
			t, _ := template.ParseFiles("./template/upload.html")
			t.Execute(w, templateParam)
			return
		}

		// 长度校验
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			log.Printf("Could not parse multipart form: %v\n", err)
			renderError(w, "CANT_PARSE_FORM", http.StatusInternalServerError)
			return
		}
		h(w, r)
	}
}

func uploadFileHandler(w http.ResponseWriter, r *http.Request) {
	// parse and validate file and post parameters
	file, fileHeader, err := r.FormFile("uploadFile")
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Get and print out file size
	fileSize := fileHeader.Size

	// validate file size
	if fileSize > maxUploadSize {
		renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	detectedFileType := http.DetectContentType(fileBytes)
	if detectedFileType != "application/x-gzip" {
		renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return
	}

	fileName := fileHeader.Filename
	newPath := filepath.Join(uploadPath, fileName)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return
	}

	if err := utils.DeCompress(newPath, uploadPath); err != nil {
		log.Print(err)
		renderError(w, "Decompress tar.gz failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("SUCCESS"))
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func shareServer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 打印来源ip及访问的文件夹/文件
		log.Printf("remote form ip:%s, uri: %s\n", r.RemoteAddr, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
