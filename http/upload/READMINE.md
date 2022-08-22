# README

## run
```
go run main.go
```

## build
go build

## reference
go标准库提供的上传下载API:
```
const maxUploadSize = 2 * 1024 * 2014 // 2 MB 
const uploadPath = "./tmp"

func main() {
    http.HandleFunc("/upload", uploadFileHandler())

    fs := http.FileServer(http.Dir(uploadPath))
    http.Handle("/files/", http.StripPrefix("/files", fs))

    log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading files.")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```
参考： https://zhuanlan.zhihu.com/p/136410759