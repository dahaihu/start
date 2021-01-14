package main

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct{}

type AddRequest struct {
	A, B int
}

type AddResponse struct {
	Res int
}

func (s *Server) Add(addReq *AddRequest, addResp *AddResponse) error {
	addResp.Res = addReq.A + addReq.B
	return nil
}

func (s *Server) Start(port string) {

	rpc.Register(s)

	rpc.HandleHTTP()
	l, e := net.Listen("tcp", port)
	if e != nil {
		fmt.Println("fatal")
	}

	http.Serve(l, nil)
}

func main() {
	s := &Server{}
	s.Start("127.0.0.1:10800")
}
