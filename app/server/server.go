package server

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/httpx"
)

type Server struct {
	Handler httpx.Handler
}

func New(handler httpx.Handler)*Server {
	return &Server{Handler: handler}
}

func(s *Server) Start() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go s.serve(conn)
	}

	
}

func(s *Server) serve(conn net.Conn) {
	defer conn.Close()

	req, err := httpx.ReadRequest(conn)
	if err != nil {
		resp := httpx.Response{
			Code: httpx.StatusBadRequest,
		}
		resp.Write(conn)
		return
	}

	response := s.Handler(req)
	response.Write(conn)

}