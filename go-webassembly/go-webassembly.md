# WebAssembly
WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine. 
Wasm is designed as a portable target for compilation of high-level languages like C/C++/Rust/Go, 
enabling deployment on the web for client and server applications.

Go 1.11 began to support WebAssembly.

## demo 
```
cd hello
GOOS=js GOARCH=wasm go build -o main.wasm
```
That will build the package and produce an executable WebAssembly module file named main.wasm.

Note that you can only compile main packages. Otherwise, you will get an object file that cannot be run in WebAssembly. 
If you have a package that you want to be able to use with WebAssembly, convert it to a main package and build a binary.

Copy the JavaScript support file:
```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

```
http-server -p 3009
(npm install -g http-server)
```
Then visit http://localhost:3009/index.html


## reference
https://github.com/golang/go/wiki/WebAssembly#getting-started