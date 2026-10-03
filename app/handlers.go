package main

import (
	"github.com/codecrafters-io/http-server-starter-go/app/httpx"
)

func echo(req *httpx.Request) *httpx.Response {
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
