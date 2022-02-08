package main

import (
	"bufio"
	"fmt"
	"net"

	"github.com/luoji_demo/tcpdemo/frame"
	"github.com/luoji_demo/tcpdemo/packet"
)

func handlePacket(framePayload []byte) (ackFramePayload []byte, err error) {
	var p packet.Packet
	p, err = packet.Decode(framePayload)
	if err != nil {
		fmt.Println("handleConn: packet decode error:", err)
		return
	}

	switch p.(type) {
	case *packet.Submit:
		submit := p.(*packet.Submit)
		fmt.Printf("recv submit: id = %s, payload=%s\n", submit.ID, string(submit.Payload))
		submitAck := &packet.SubmitAck{
			ID:     submit.ID,
			Result: 0,
		}
		ackFramePayload, err = packet.Encode(submitAck)
		if err != nil {
			fmt.Println("handleConn: packet encode error:", err)
			return nil, err
		}
		return ackFramePayload, nil
	default:
		return nil, fmt.Errorf("unknown packet type")
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	frameCodec := frame.NewMyFrameCodec()

	/* 我们的优化目标是降低 net.Conn 的 Write 和 Read 的频率
	将 net.Conn 改为 rbuf 后，frameCodec.Decode 中的每次网络读取实际调用的都是 bufio.Reader 的 Read 方法。
	bufio.Reader.Read 方法内部，每次从 net.Conn 尝试读取其内部缓存大小的数据，而不是用户传入的希望读取的数据大小。
	这些数据缓存在内存中，这样，后续的 Read 就可以直接从内存中得到数据，而不是每次都要从 net.Conn 读取，从而降低 Syscall 调用的频率。
	*/
	rbuf := bufio.NewReader(c)
	wbuf := bufio.NewWriter(c)

	for {
		// decode the frame to get the payload
		// 用buf取代net.Conn, framePayload, err := frameCodec.Decode(c)
		framePayload, err := frameCodec.Decode(rbuf)
		if err != nil {
			fmt.Println("handleConn: frame decode error:", err)
			return
		}

		// do something with the packet
		ackFramePayload, err := handlePacket(framePayload)
		if err != nil {
			fmt.Println("handleConn: handle packet error:", err)
			return
		}

		// write ack frame to the connection
		err = frameCodec.Encode(wbuf, ackFramePayload)
		if err != nil {
			fmt.Println("handleConn: frame encode error:", err)
			return
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle the new connection.
		go handleConn(c)
	}
}
