package main

import (
	"fmt"
	"log"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
	"github.com/codecrafters-io/http-server-starter-go/app/httpx"
	"github.com/codecrafters-io/http-server-starter-go/app/middleare"
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	cfg := config.New()
	handlers := NewHandlers(cfg)

	router := httpx.NewRouter()
	router.AddRoute(httpx.GET, "/", func(req *httpx.Request) *httpx.Response {
		return httpx.Text(httpx.StatusOK, "")
	})
	router.AddRoute(httpx.GET, "/echo/{str}", echo)
	router.AddRoute(httpx.GET, "/user-agent", user_agent)
	router.AddRoute(httpx.GET, "/files/{filename}", handlers.file)
	router.AddRoute(httpx.POST, "/files/{filename}", handlers.createFile)

	s := server.New(middleare.Compress(router.Route))
	err := s.Start()
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}
