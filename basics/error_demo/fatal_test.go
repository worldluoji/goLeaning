package error_demo

import (
	"log"
	"testing"
)

func foo() {
	log.Println("foo")
}

func fatal() {
	log.Fatal("fatal")
}

// log.Fatal执行后，defer将不会执行
func TestFatal(t *testing.T) {
	defer func() {
		log.Println("defer")
	}()
	go foo()
	go fatal()
	go foo()
	select {}
}
