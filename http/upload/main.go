package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	maxUploadSize = 3 * 1024 * 2014 // 3 MB
	uploadPath    = "D:\\temp"
)

type HTTPHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func handler(h http.HandlerFunc, decorators ...HTTPHandlerDecorator) http.HandlerFunc {
	for i := range decorators {
		d := decorators[len(decorators)-i-1] // in reverse
		h = d(h)
	}
	return h
}

func main() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/upload", handler(uploadFileHandler, preChecker))
	serverMux.HandleFunc("/view", viewHandler)

	log.Print("Server started on por 8080, use /upload for uploading files")
	server := &http.Server{
		Addr:    ":8080",
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
			t.Execute(w, nil)
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
	// log.Printf("File size (bytes): %v\n", fileSize)
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
	// log.Print("File type " + detectedFileType)
	if detectedFileType != "application/x-gzip" {
		renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return
	}

	fileName := fileHeader.Filename
	newPath := filepath.Join(uploadPath, fileName)
	log.Printf("FileType: %s, File: %s\n", detectedFileType, newPath)

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
	w.Write([]byte("SUCCESS"))
}

func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./template/upload.html")
	t.Execute(w, nil)
}
