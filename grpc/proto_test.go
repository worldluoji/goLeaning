package gprc

import (
	"testing"

	"github.com/golang/protobuf/proto"
	tc "github.com/luoji_demo/grpc/grpctest"
)

func TestProto(t *testing.T) {
	p := &tc.Person{
		Name: "xiaoming",
		Age:  10,
	}
	bytes, err := proto.Marshal(p)
	if err != nil {
		t.Log("fail to encode by protobuf")
		panic(err)
	}
	t.Log(bytes)
	up := &tc.Person{}
	err = proto.Unmarshal(bytes, up)
	if err != nil {
		t.Log("fail to decode by protobuf")
		panic(err)
	}

	t.Log(up)
}
