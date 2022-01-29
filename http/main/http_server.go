package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
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
		for k, v := range r.Header {
			for _, s := range v {
				w.Header().Add(k, s)
			}
		}

		version := os.Getenv("JAVA_HOME")
		if version != "" {
			w.Header().Set("VERSION", version)
		} else {
			w.Header().Set("VERSION", "UnKnown")
		}

		w.WriteHeader(200)

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
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler(echo, WithServerHeader, WithAuthCookie))
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	server := &http.Server{
		Addr:    ":8088",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
