package server

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/httpx"
)

type Server struct {
	Handler httpx.Handler
}

func New(handler httpx.Handler)*Server {
	return &Server{Handler: handler}
}

func(s *Server) Start() error {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		return err
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go s.serve(conn)
	}
}

func(s *Server) serve(conn net.Conn) {
	for {
		req, err := httpx.ReadRequest(conn)
		if err != nil {
			fmt.Println(err.Error())
			resp := httpx.Response{
				Code: httpx.StatusBadRequest,
			}
			err := resp.Write(conn)
			if err != nil {
				fmt.Println(err.Error())
			}
			return
		}

		response := s.Handler(req)
		if header := req.Header(httpx.Connection); header != nil && header.Value == "close" {
			response.AddHeader(httpx.Connection, "close")
			response.Write(conn)
			conn.Close()
			break
		
		}

		response.Write(conn)
	}
}