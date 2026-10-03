package httpx

import (
	"io"
	"strconv"
	"strings"
)

type Response struct {
	Code    int
	Headers []*Header
	Body    []byte
}

const (
	StatusOK         = 200
	StatusCreated    = 201
	StatusBadRequest = 400
	StatusNotFound   = 404
)

func (r *Response) Marshal() []byte {
	response := "HTTP/1.1 "

	response += strconv.Itoa(r.Code)
	response += " " + getStatusByCode(r.Code) + "\r\n"

	for _, header := range r.Headers {
		if strings.EqualFold(header.Name, ContentLength) {
			continue
		}
		response += header.Name + ": " + header.Value + "\r\n"
	}
	response += ContentLength + ": " + strconv.Itoa(len(r.Body)) + "\r\n"
	response += "\r\n"

	response += string(r.Body)

	return []byte(response)
}

func getStatusByCode(code int) string {
	switch code {
	case StatusOK:
		return "OK"
	case StatusBadRequest:
		return "Bad Request"
	case StatusNotFound:
		return "Not Found"
	case StatusCreated:
		return "Created"
	}

	return "Ok"
}

func (r Response) Write(writer io.Writer) error {
	_, err := writer.Write(r.Marshal())
	return err
}

func (r *Response) AddHeader(name string, value string) {
	r.Headers = append(r.Headers, &Header{Name: name, Value: value})
}

func (r *Response) Header(name string) *Header {
	for _, header := range r.Headers {
		if strings.EqualFold(header.Name, name) {
			return header
		}
	}

	return nil
}

func Text(code int, body string) *Response {
	return NewResponse(
		code,
		map[string]string{"Content-Type": "text/plain"},
		body,
	)
}

func NotFound() *Response {
	return NewResponse(
		StatusNotFound,
		map[string]string{"Content-Type": "text/plain"},
		"",
	)
}

func NewResponse(code int, headers map[string]string, body string) *Response {
	var hs []*Header

	for name, vlaue := range headers {
		hs = append(hs, &Header{Name: name, Value: vlaue})
	}

	return &Response{
		Code:    code,
		Headers: hs,
		Body:    []byte(body),
	}
}
