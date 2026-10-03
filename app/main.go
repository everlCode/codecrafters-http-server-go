package main

import (
	"fmt"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/httpx"
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	router := httpx.NewRouter()
	router.AddRoute("/", func(req *httpx.Request) *httpx.Response {
		return httpx.Text(httpx.StatusOK, "")
	})
	router.AddRoute("/echo/{str}", echo)
	router.AddRoute("/user-agent", user_agent)

	server := server.New(router.Route)
	server.Start()
}
