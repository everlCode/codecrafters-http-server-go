package parser

import (
	"bufio"
	"log"
	"net"
	"strings"
)

type HttpParser struct {

}

type HttpRequest struct {
	Method string
	Uri string
	Version string
	Headers []*Header
}

type Header struct {

}

func ReadRequest(conn net.Conn) *HttpRequest {
	r := bufio.NewReader(conn)
	line, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err.Error())
	}
	parts := strings.Split(string(line) , " ")
	if len(parts) < 3 {
		log.Fatal("Некоректный запрос!")
	}


	return &HttpRequest{
		Method: parts[0],
		Uri: parts[1],
		Version: parts[2],
	}
}