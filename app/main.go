package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
	"github.com/codecrafters-io/http-server-starter-go/app/router"
	"github.com/codecrafters-io/http-server-starter-go/app/writer"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// TODO: Uncomment the code below to pass the first stage
	
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	
	conn, err := l.Accept()
	defer conn.Close()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	req := parser.ReadRequest(conn)
	router := router.New(
		map[string]func(*parser.HttpRequest) *parser.HttpResponse{"/": func(req *parser.HttpRequest) *parser.HttpResponse {
			return &parser.HttpResponse{
				Code: 200,
			}
		}},
	)
	response := router.Route(req)

	writer := writer.New(conn)
	writer.Write(response)
}
