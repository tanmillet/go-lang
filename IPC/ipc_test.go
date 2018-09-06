package IPC

import (
	"testing"
	"fmt"
)

type EchoServer struct {
}

func NewEchoServer() *EchoServer {
	return &EchoServer{}
}

func (server *EchoServer) Handle(methon, request string) *Response {
	return &Response{}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(NewEchoServer())

	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)
	resp1, _ := client1.Call("POST", "From Client1")
	fmt.Println(resp1)
	resp2, _ := client2.Call("GET", "From Client1")
	fmt.Println(resp2)

	defer client1.Close()
	defer client2.Close()
}
