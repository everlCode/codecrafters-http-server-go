package parser

import "strconv"

type HttpResponse struct {
	Code int
	Method string
	Uri string
	Version string
}

func (hr *HttpResponse) Marshal()[]byte {
	response := "HTTP/1.1 "

	response += strconv.Itoa(hr.Code)
	response += " " + getStatusByCode(hr.Code) + "\r\n\r\n"


	return []byte(response)
}

func getStatusByCode(code int) string {
	switch code {
	case 200:
		return "OK"
	case 404:
		return "Not Found"
	}

	return "Ok"
}