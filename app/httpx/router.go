package httpx

import "strings"

type Router struct {
	routes []*route
}

type route struct {
	Method string
	segments []*segment
	Handler  Handler
}

type segment struct {
	Name    string
	isParam bool
}

const (
	GET = "GET"
	POST = "POST"
)

type Handler func(*Request) *Response

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Route(req *Request) *Response {
	route, params := r.match(req)
	if route == nil {
		return NotFound()
	}
	req.Params = params

	return route.Handler(req)
}

func (r *Router) AddRoute(method string, path string, handler Handler) {
	route := route{
		Method: method,
		Handler: handler,
	}

	for _, part := range strings.Split(path, "/") {
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			route.segments = append(route.segments, &segment{Name: strings.Trim(part, "{}"), isParam: true})
		} else {
			route.segments = append(route.segments, &segment{Name: part})
		}
	}

	r.routes = append(r.routes, &route)
}

func (r *Router) match(req *Request) (*route, map[string]string) {
	pathParts := strings.Split(req.URI, "/")
	for _, route := range r.routes {
		if route.Method != req.Method {
			continue
		}
		if len(route.segments) != len(pathParts) {
			continue
		}
		
		ok, params := route.match(pathParts)
		if ok {
			return route, params
		}
	}

	return nil, map[string]string{}
}

func (r *route) match(parts []string) (bool, map[string]string) {
	params := make(map[string]string)
	for k, part := range parts {
		if part != r.segments[k].Name && !r.segments[k].isParam {
			return false, nil
		}
		if r.segments[k].isParam {
			params[r.segments[k].Name] = part
		}
	}

	return true, params
}