package main

import (
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
)

func echo(req *http.Request) *http.Response {
	param, ok := req.Params["str"]
	if !ok {
		return &http.Response{
			Code: http.StatusNotFound,
		}
	}

	return http.NewResponse(
		http.StatusOK,
		map[string]string{"Content-Type": "text/plain", "Content-Length": strconv.Itoa(len(param))},
		param,
	)
}


func user_agent(req *http.Request) *http.Response {
	header := req.Header(http.HeaderUserAgent)
	if header == nil {
		return &http.Response{
			Code: http.StatusNotFound,
		}
	}

	return http.NewResponse(
		http.StatusOK,
		map[string]string{"Content-Type": "text/plain", "Content-Length": strconv.Itoa(len(header.Value))},
		header.Value,
	)
}
