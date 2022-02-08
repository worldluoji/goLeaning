package frame

import (
	"encoding/binary"
	"errors"
	"io"
)

type FramePayload []byte

// Decode就是指字节流 -> Frame，Encode是指Frame -> 字节流
type StreamFrameCodec interface {
	Encode(io.Writer, FramePayload) error   // data -> frame，并写入io.Writer
	Decode(io.Reader) (FramePayload, error) // 从io.Reader中提取frame payload，并返回给上层
}

var ErrShortWrite = errors.New("short write")
var ErrShortRead = errors.New("short read")

type myFrameCodec struct{}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}

func (p *myFrameCodec) Encode(w io.Writer, framePayload FramePayload) error {
	var f = framePayload
	// 长度字段为uint32,占4字节
	var totalLen int32 = int32(len(framePayload)) + 4

	// 网络字节序使用大端字节序（BigEndian），因此无论是 Encode 还是 Decode，我们都是用 binary.BigEndian；
	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}

	n, err := w.Write([]byte(f)) // write the frame payload to outbound stream
	if err != nil {
		return err
	}

	if n != len(framePayload) {
		return ErrShortWrite
	}

	return nil
}

func (p *myFrameCodec) Decode(r io.Reader) (FramePayload, error) {
	var totalLen int32

	// binary.Read 或 Write 会根据参数的宽度，读取或写入对应的字节个数的字节，这里 totalLen 使用 int32，那么 Read 或 Write 只会操作数据流中的 4 个字节
	err := binary.Read(r, binary.BigEndian, &totalLen)
	if err != nil {
		return nil, err
	}

	if totalLen <= 4 {
		return nil, errors.New("read data length error")
	}

	buf := make([]byte, totalLen-4)
	// io.ReadFull 一般会读满你所需的字节数，除非遇到 EOF 或 ErrUnexpectedEOF
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	if n != int(totalLen-4) {
		return nil, ErrShortRead
	}

	return FramePayload(buf), nil
}
