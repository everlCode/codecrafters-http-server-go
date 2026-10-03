package http

import (
	"bufio"
	"errors"
	"io"
	"strings"
)

type Request struct {
	Method  string
	URI     string
	Version string
	Headers []*Header
	Params map[string]string
}

type Header struct {
	Name string
	Value string
}

const (
	HeaderUserAgent = "User-Agent"
)

func ReadRequest(reader io.Reader) (*Request, error) {
	r := bufio.NewReader(reader)
	line, _, err := r.ReadLine()
	if err != nil {
		return nil, err
	}
	parts := strings.Split(string(line), " ")
	if len(parts) < 3 {
		return nil, errors.New("Неккоректный запрос!")

	}
	var headers []*Header
	for {
		header, _, err := r.ReadLine()
		if err != nil {
			 return nil, err
		}
		if len(header) == 0 {
			break
		}
		headerParts := headerParts(string(header))
		if len(headerParts) < 2 {
			continue
		}
		headers = append(headers, &Header{Name: headerParts[0], Value: headerParts[1]})
	}

	return &Request{
		Method:  parts[0],
		URI:     parts[1],
		Version: parts[2],
		Headers: headers,
	}, nil
}

func headerParts(header string) []string {
	before, after, ok := strings.Cut(header, ":")
	if !ok {
		return []string{}
	}

	return []string{strings.Trim(before, " "), strings.Trim(after, " ")}
}

func (r *Request) Header(name string) *Header {
	for _, header := range r.Headers {
		if header.Name == name {
			return header
		}
	}

	return nil
}
