package iotest

import (
	"bufio"
	"io"
	"os"
	"strings"
	"testing"
)

func CountLines(r io.Reader) (int, error) {
	scanner := bufio.NewScanner(r)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return lines, scanner.Err()
}

func TestReadLine(t *testing.T) {
	r := strings.NewReader("cat\nfish\ndog")
	lines, err := CountLines(r)
	if err != nil {
		t.Log(err)
	} else {
		t.Log(lines)
	}
}

func TestReadFile(t *testing.T) {
	f, err := os.Open("words.txt")
	if err != nil {
		t.Log("fail to open file")
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		t.Log(scanner.Text())
	}
}
