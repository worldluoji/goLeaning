package main

import (
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// read from the connection
		var buf = make([]byte, 128)

		// SetReadDeadline 和 SetWriteDeadline 一样，只要我们通过 SetWriteDeadline 设置了超时，那无论后续方法是否成功，如果不重新设置写超时或取消写超时，后续对 Socket 的写操作都将以超时失败告终。
		c.SetReadDeadline(time.Now().Add(time.Second))

		/*
		* 即便是 Read 操作，也是有 lock 保护的。多个 Goroutine 对同一conn的并发读，
		* 不会出现读出内容重叠的情况，但就像前面讲并发读的必要性时说的那样，
		* 一旦采用了不恰当长度的切片作为 buf，很可能读出不完整的业务包，这反倒会带来业务上的处理难度。
		* 所以服务器端一个goutine来处理一个连接，而消费方一个conn发送一个完整的数据包，这样是好的处理方案
		 */
		n, err := c.Read(buf)
		if err != nil {
			log.Printf("conn read %d bytes,  error: %s", n, err)
			if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
				// 进行其他业务逻辑的处理
				continue
			}
			return
		}
		log.Printf("read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}
