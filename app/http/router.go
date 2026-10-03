package http

import "strings"

type Router struct {
	Routes []*Route
}

type Route struct {
	Name     string
	Segments []*segment
	Handler  Handler
}

type segment struct {
	Name    string
	isParam bool
}

type Handler func(*Request) *Response

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Route(req *Request) *Response {
	route, params := r.GetRoute(req.URI)
	if route == nil {
		return &Response{
			Code: StatusNotFound,
		}
	}
	req.Params = params

	return route.Handler(req)
}

func (r *Router) AddRoute(path string, handler Handler) {
	route := Route{Name: path, Handler: handler}

	for _, part := range strings.Split(path, "/") {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			route.Segments = append(route.Segments, &segment{Name: strings.Trim(part, "{}"), isParam: true})
		} else {
			route.Segments = append(route.Segments, &segment{Name: part})
		}
	}

	r.Routes = append(r.Routes, &route)
}

func (r *Router) GetRoute(path string) (*Route, map[string]string) {
	pathParts := strings.Split(path, "/")
	for _, route := range r.Routes {
		if len(route.Segments) != len(pathParts) {
			continue
		}

		var isEqual bool = true
		params := make(map[string]string)
		for k, part := range pathParts {
			if part != route.Segments[k].Name && !route.Segments[k].isParam {
				isEqual = false
				break
			}
			if route.Segments[k].isParam {
				params[route.Segments[k].Name] = part
			}

		}
		if isEqual {
			return route, params
		}
	}

	return nil, map[string]string{}
}