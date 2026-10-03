package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
	"github.com/codecrafters-io/http-server-starter-go/app/httpx"
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	fmt.Println("Logs from your program will appear here!")

	cfg := config.New()
	handlers := NewHandlers(cfg)

	router := httpx.NewRouter()
	router.AddRoute("/", func(req *httpx.Request) *httpx.Response {
		return httpx.Text(httpx.StatusOK, "")
	})
	router.AddRoute("/echo/{str}", handlers.echo)
	router.AddRoute("/user-agent", handlers.user_agent)
	router.AddRoute("/files/{filename}", handlers.file)

	s := server.New(router.Route)
	err := s.Start()
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}
