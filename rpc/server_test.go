package main

import (
	"fmt"
	"net/rpc"
	"testing"
)

func TestRpcServer(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "localhost:10800")
	if err != nil {
		fmt.Printf("error %s", err)
		return
	}
	req := &AddRequest{20, 20}
	resp := &AddResponse{}
	err = client.Call("Server.Add", req, resp)
	if err != nil {
		fmt.Println("err is ", err)
	}
	fmt.Println("resp is ", resp)
}