package main

import (
	"strconv"

	"github.com/codecrafters-io/http-server-starter-go/app/http"
)

func echo(req *http.Request) *http.Response {
	if len(req.Route.Keys) == 0 {
		return &http.Response{
			Code: http.StatusNotFound,
		}
	}

	return http.HTTPresponse(
		http.StatusOK,
		map[string]string{"Content-Type": "text/plain", "Content-Length": strconv.Itoa(len(req.Route.Keys["str"]))},
		req.Route.Keys["str"],
	)
}