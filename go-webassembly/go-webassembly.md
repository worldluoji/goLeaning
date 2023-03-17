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

## Executing WebAssembly with Node.js
First, make sure Node is installed and in your PATH.

then add $(go env GOROOT)/misc/wasm to your PATH
```
export PATH="$PATH:$(go env GOROOT)/misc/wasm"
```
This will allow go run and go test find go_js_wasm_exec in a `PATH search and use it to just work for js/wasm.
```
GOOS=js GOARCH=wasm go run .
```

go_js_wasm_exec is a wrapper that allows running Go Wasm binaries in Node. 
By default, it may be found in the misc/wasm directory of your Go installation.

If you’d rather not add anything to your PATH, you may also set the -exec flag to the location of go_js_wasm_exec 
when you execute go run or go test manually.
```
$ GOOS=js GOARCH=wasm go run -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" .
Hello, WebAssembly!

$ GOOS=js GOARCH=wasm go test -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec"
PASS
ok  	example.org/my/pkg	0.800s

$(go env GOROOT)/misc/wasm/go_js_wasm_exec ./main.wasm
Hello, WebAssembly!
```

## Running tests in the browser
```
go install github.com/agnivade/wasmbrowsertest@latest

// $GOPATH is $(go env GOPATH)

$ mv $GOPATH/bin/wasmbrowsertest $GOPATH/bin/go_js_wasm_exec
$ export PATH="$PATH:$GOPATH/bin"
$ GOOS=js GOARCH=wasm go test
PASS
ok  	example.org/my/pkg	0.800s
```
It automates the job of spinning up a webserver and uses headless Chrome to run the tests inside it 
and relays the logs to your console.

<br>

## reference
https://github.com/golang/go/wiki/WebAssembly#getting-started