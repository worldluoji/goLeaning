package other_demo

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

func TestFatal(t *testing.T) {
	defer func() {
		log.Println("defer")
	}()
	go foo()
	go fatal()
	go foo()
	select {}
}
