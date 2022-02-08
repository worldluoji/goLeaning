package main

import (
	"log"
	"net"
	"time"
)

func main() {
	//conn, err := net.Dial("tcp", "localhost:8888")
	conn, err := net.DialTimeout("tcp", "localhost:8888", 2*time.Second)
	if err != nil {
		log.Println("Dial error:", err)
		return
	}

	defer conn.Close()
	log.Println("Dial ok")

	conn.SetWriteDeadline(time.Now().Add(time.Microsecond * 100))
	/*
	* 每次 Write 操作都是受 lock 保护，直到这次数据全部写完才会解锁。因此，在应用层面，要想保证多个 Goroutine 在一个conn上 write 操作是安全的，
	* 需要一次 write 操作完整地写入一个“业务包”。一旦将业务包的写入拆分为多次 write，那也无法保证某个 Goroutine 的某“业务包”数据在conn发送的连续性。
	 */
	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println("Write error:", err)
		return
	}

	log.Println("Write ", n, " bytes to server")

}
