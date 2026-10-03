package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/codecrafters-io/http-server-starter-go/app/config"
	"github.com/codecrafters-io/http-server-starter-go/app/httpx"
)

type Handlers struct {
	Config *config.Config
}

func NewHandlers(cfg *config.Config) *Handlers {
	return &Handlers{Config: cfg}
}

func echo(req *httpx.Request) *httpx.Response {
	param, ok := req.Params["str"]
	if !ok {
		return httpx.NotFound()
	}

	return httpx.Text(
		httpx.StatusOK,
		param,
	)
}

func user_agent(req *httpx.Request) *httpx.Response {
	header := req.Header(httpx.HeaderUserAgent)
	if header == nil {
		return &httpx.Response{
			Code: httpx.StatusBadRequest,
		}
	}

	return httpx.Text(
		httpx.StatusOK,
		header.Value,
	)
}

func (h *Handlers) file(req *httpx.Request) *httpx.Response {
	filename, ok := req.Params["filename"]
	if !ok {
		return httpx.NotFound()
	}

	file, err := os.ReadFile(filepath.Join(h.Config.Dirname, filename))
	if err != nil {
		wd, _ := os.Getwd()
		log.Printf("open %q (cwd=%q): %v", file, wd, err)
		return httpx.NotFound()
	}

	return httpx.NewResponse(
		httpx.StatusOK,
		map[string]string{httpx.ContentType: "application/octet-stream"},
		string(file),
	)
}

func (h *Handlers) createFile(req *httpx.Request) *httpx.Response {
	filename, ok := req.Params["filename"]
	if !ok {
		return httpx.NotFound()
	}

	file, err := os.Create(filepath.Join(h.Config.Dirname, filename))
	defer file.Close()
	if err != nil {
		return httpx.NotFound()
	}
	_, err = file.Write(req.Body)
	if err != nil {
		return httpx.NewResponse(
			httpx.StatusBadRequest,
			map[string]string{httpx.ContentType: "application/octet-stream"},
			"",
		)
	}

	return httpx.NewResponse(
		httpx.StatusCreated,
		map[string]string{httpx.ContentType: "application/octet-stream"},
		"",
	)
}
