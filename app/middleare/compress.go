package middleare

import (
	"bytes"
	"compress/gzip"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/httpx"
)

type encoder func([]byte) ([]byte, error)

var encoders = map[string]encoder{
	"gzip": gzipCompress,
}


func Compress(next httpx.Handler) httpx.Handler {
	return func(req *httpx.Request) *httpx.Response {
		resp := next(req)

		var header *httpx.Header
		if header = req.Header(httpx.AcceptEncoding);header == nil {
			return resp
		}
		compressMethod := strings.ToLower(header.Value)
		encodeHandler, ok := encoders[compressMethod]
		if !ok {
			return resp
		}
		compressedBody, err := encodeHandler(resp.Body)
		if err != nil {
			return resp
		}
		resp.Body = compressedBody
		resp.AddHeader(httpx.ContentEncoding, compressMethod)

		return resp
	}
}

func gzipCompress(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	_, err := writer.Write(input);if err != nil {
		return nil, err
	}
	writer.Close()


	return buf.Bytes(), nil
}