package router

import "github.com/codecrafters-io/http-server-starter-go/app/parser"

type Router struct {
	Routes map[string]func(*parser.HttpRequest) *parser.HttpResponse
}

func New(routes map[string]func(*parser.HttpRequest) *parser.HttpResponse)  *Router {
	return &Router{
		Routes: routes,
	}
}

func (r *Router) Route(req *parser.HttpRequest) *parser.HttpResponse {
	route, ok := r.Routes[req.Uri]
	if !ok {
		return &parser.HttpResponse{
			Code: 404,
		}
	}

	return route(req)
}