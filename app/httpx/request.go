package httpx

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

type Request struct {
	Method  string
	URI     string
	Version string
	Headers []*Header
	Body    []byte
	Params  map[string]string
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
	var headers []*Header
	var contentLenght int
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
		if headerParts[0] == ContentLength {
			contentLenght, err = strconv.Atoi(headerParts[1])
			if err != nil {
				return nil, err
			}
		}
		headers = append(headers, &Header{Name: headerParts[0], Value: headerParts[1]})
	}
	buff := make([]byte, contentLenght)
	_, err = r.Read(buff)
	if err != nil {
		return nil, err
	}

	return &Request{
		Method:  strings.ToLower(parts[0]),
		URI:     parts[1],
		Version: parts[2],
		Headers: headers,
		Body:    buff,
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
		if strings.EqualFold(header.Name, name) {
			return header
		}
	}

	return nil
}
