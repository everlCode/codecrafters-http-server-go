package httpx

import (
	"io"
	"strconv"
)

type Response struct {
	Code    int
	Headers []*Header
	Body    string
}

const (
	StatusOK         = 200
	StatusBadRequest = 400
	StatusNotFound   = 404
)

func (r *Response) Marshal() []byte {
	response := "HTTP/1.1 "

	response += strconv.Itoa(r.Code)
	response += " " + getStatusByCode(r.Code) + "\r\n"

	for _, header := range r.Headers {
		response += header.Name + ": " + header.Value + "\r\n"
	}
	response += "\r\n"
	response += r.Body

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
	}

	return "Ok"
}

func (hr Response) Write(writer io.Writer) error {
	_, err := writer.Write(hr.Marshal())
	return err
}

func Text(code int, body string) *Response {
	return NewResponse(
		code,
		map[string]string{"Content-Type": "text/plain", "Content-Length": strconv.Itoa(len(body))},
		body,
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
		Body:    body,
	}
}
