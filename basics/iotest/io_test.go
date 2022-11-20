package iotest

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"reflect"
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

/*
*  io.Writer 接口类型表示数据写入的目的地，既可以支持向磁盘写入，也可以支持向网络存储写入，
*  并支持任何实现了 Write 方法的写入行为，这让 Save 函数的扩展性得到了质的提升。
*
*  用 io.Writer 接口类型替换掉了 *os.File。这样一来，新版 Save 的设计就符合了接口分离原则，
*  因为 io.Writer 仅包含一个 Write 方法，而且这个方法恰恰是 Save 唯一需要的方法。
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
*/
func Save(w io.Writer, data []byte) error {
	if _, err := w.Write(data); err != nil {
		return err
	}

	return nil
}

func TestSave(t *testing.T) {
	b := make([]byte, 0, 128)
	buf := bytes.NewBuffer(b)
	data := []byte("hello, golang")
	err := Save(buf, data)
	if err != nil {
		t.Errorf("want nil, actual %s", err.Error())
	}

	saved := buf.Bytes()
	if !reflect.DeepEqual(saved, data) {
		t.Errorf("want %s, actual %s", string(data), string(saved))
	} else {
		t.Log("buf and data equals")
	}

}
