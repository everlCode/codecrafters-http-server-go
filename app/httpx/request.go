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
	method, uri, version, err := readRequestLine(r)
	if err != nil {
		return nil, err
	}
	req := &Request{
		Method:  method,
		URI:     uri,
		Version: version,
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
	req.Headers = headers
	var contentLenght int
	contentLengthHeader := req.Header(ContentLength);if contentLengthHeader != nil {
		contentLenght, _ = strconv.Atoi(contentLengthHeader.Value)
	}
	
	if contentLenght > 0 {
		buff := make([]byte, contentLenght)
		_, err = r.Read(buff)
		if err != nil {
			return nil, err
		}
		req.Body = buff
	}
	
	return req, nil
}

func readRequestLine(r *bufio.Reader) (string, string, string, error) {
	line, _, err := r.ReadLine()
	if err != nil {
		return "", "", "", err
	}
	parts := strings.Split(string(line), " ")
	if len(parts) < 3 {
		return "", "", "", errors.New("Неккоректный запрос!")

	}

	return strings.ToUpper(parts[0]), parts[1], parts[2], nil
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
