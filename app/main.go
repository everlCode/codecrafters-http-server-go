package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
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

	router := http.NewRouter()
	router.AddRoute("/", func(req *http.Request) *http.Response {
		return &http.Response{
			Code: http.StatusOK,
		}
	})
	router.AddRoute("/echo/{str}", echo)
	router.AddRoute("/user-agent", user_agent)

	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	req, err := http.ReadRequest(conn)
	if err != nil {
		resp := http.Response{
			Code: http.StatusBadRequest,
		}
		resp.Write(conn)
		return
	}

	response := router.Route(req)
	response.Write(conn)
}
