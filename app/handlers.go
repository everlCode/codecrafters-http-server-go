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

func (h *Handlers) echo(req *httpx.Request) *httpx.Response {
	param, ok := req.Params["str"]
	if !ok {
		return &httpx.Response{
			Code: httpx.StatusNotFound,
		}
	}

	return httpx.Text(
		httpx.StatusOK,
		param,
	)
}


func (h *Handlers) user_agent(req *httpx.Request) *httpx.Response {
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
		return &httpx.Response{
			Code: httpx.StatusNotFound,
		}
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
