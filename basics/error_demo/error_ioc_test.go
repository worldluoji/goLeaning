package error_demo

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"testing"
)

// 长度不够，少一个Weight，这里只有11个字节，但是Person结构体，有12个字节， 1字节=8bit
var b = []byte{0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c}
var r = bytes.NewReader(b)

// 将error定义在Person对象中，实际是一种反转控制，当binary Read发生错误时，就会记录err信息，代码会更简洁
type Person struct {
	Name   [10]byte
	Age    uint8
	Weight uint8
	err    error
}

func (p *Person) read(data interface{}) {
	if p.err == nil {
		// 按照大端序，从r中读取内容，写到data中
		p.err = binary.Read(r, binary.BigEndian, data)
	}
}

func (p *Person) ReadName() *Person {
	p.read(&p.Name)
	return p
}
func (p *Person) ReadAge() *Person {
	p.read(&p.Age)
	return p
}
func (p *Person) ReadWeight() *Person {
	p.read(&p.Weight)
	return p
}
func (p *Person) Print() *Person {
	if p.err == nil {
		fmt.Printf("Name=%s, Age=%d, Weight=%d\n", p.Name, p.Age, p.Weight)
	}
	return p
}

func TestErrorIOC(t *testing.T) {
	p := Person{}
	p.ReadName().ReadAge().ReadWeight().Print()
	t.Log(p.err) // EOF 错误
}

/*
我们还可以封装自己的Error, 提供其它信息，例如：
// $GOROOT/src/net/net.go
type Error interface {
    error
    Timeout() bool
    Temporary() bool
}


// $GOROOT/src/net/http/server.go
func (srv *Server) Serve(l net.Listener) error {
    ... ...
    for {
        rw, e := l.Accept()
        if e != nil {
            select {
            case <-srv.getDoneChan():
                return ErrServerClosed
            default:
            }
            if ne, ok := e.(net.Error); ok && ne.Temporary() {
                // 注：这里对临时性(temporary)错误进行处理
                ... ...
                time.Sleep(tempDelay)
                continue
            }
            return e
        }
        ...
    }
    ... ...
}
*/
