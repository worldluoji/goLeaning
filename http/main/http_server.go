package main

import (
	"fmt"
	"log"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Receive request %s from %s", r.URL.Path, r.RemoteAddr)
	fmt.Fprintf(w, "hello "+r.URL.Path)
}

// WithServerHeader decorator example
func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("-->WithServerHeader()")
		w.Header().Set("Server", "EchoSever 0.0.1")
		h(w, r)
	}
}

// WithAuthCookie example
func WithAuthCookie(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("--->WithAuthCookie()")
		cookie := &http.Cookie{Name: "Auth", Value: "Pass", Path: "/"}
		http.SetCookie(w, cookie)
		h(w, r)
	}
}

// HTTPHandlerDecorator for http.HandlerFunc
type HTTPHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func handler(h http.HandlerFunc, decorators ...HTTPHandlerDecorator) http.HandlerFunc {
	for i := range decorators {
		d := decorators[len(decorators)-i-1] // in reverse
		h = d(h)
	}
	return h
}

func main() {
	http.HandleFunc("/v1/echo", handler(echo, WithServerHeader, WithAuthCookie))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
