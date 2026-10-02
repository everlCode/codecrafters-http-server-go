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
	Route *Route
}

type Header struct {
	Name string
	Value string
}

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

	return &Request{
		Method:  parts[0],
		URI:     parts[1],
		Version: parts[2],
	}, nil
}
