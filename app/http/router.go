package http

import "strings"

type Router struct {
	Routes map[string]Handler
}

type Route struct {
	Keys map[string]string
	Handler Handler
}

type Handler func(*Request) *Response

func NewRouter(routes map[string]Handler) *Router {
	return &Router{
		Routes: routes,
	}
}

func (r *Router) Route(req *Request) *Response {
	route := r.GetRoute(req.URI)
	if route == nil {
		return &Response{
			Code: StatusNotFound,
		}
	}
	req.Route = route

	return route.Handler(req)
}

func (r *Router) AddRoute(path string, handler Handler) {
	r.Routes[path] = handler
}

func (r *Router) GetRoute(path string) *Route {
	pathParts := strings.Split(path, "/")
	for route, h := range r.Routes {
		r := Route{}
		r.Handler = h
		keys := make(map[string]string)
		routeParts := strings.Split(route, "/")
		if len(routeParts) != len(pathParts) {
			continue
		}

		var isEqual bool = true
		for k, part := range pathParts {
			if part != routeParts[k] && !isPartKey(routeParts[k]) {
				isEqual = false
				break
			}
			if isPartKey(routeParts[k]) {
				keys[clearKey(routeParts[k])] = part
			}
		}
		if isEqual {
			r.Keys = keys
			return &r
		}
		
	}


	return nil
}

func isPartKey(path string) bool {
	return strings.HasPrefix(path, "{")
}

func clearKey(path string) string {
	str := strings.Replace(path, "{", "", 1)
	return strings.Replace(str, "}", "", 1)
}



