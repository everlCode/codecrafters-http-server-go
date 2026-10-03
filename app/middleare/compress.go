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
		method, encodeHandler := selectEncoder(header.Value);
		if encodeHandler == nil {
			return resp
		}
		compressedBody, err := encodeHandler(resp.Body)
		if err != nil {
			return resp
		}
		resp.Body = compressedBody
		resp.AddHeader(httpx.ContentEncoding, method)

		return resp
	}
}

func gzipCompress(input []byte) ([]byte, error) {
	if len(input) == 0 {
		return input, nil
	}
	var buf bytes.Buffer
	writer := gzip.NewWriter(&buf)
	if _, err := writer.Write(input); err != nil {
		return nil, err
	}
	err := writer.Close();if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func selectEncoder(methods string) (string, encoder) {
	parts := strings.Split(methods, ",")

	for _, method := range parts {
		method = strings.ToLower(strings.TrimSpace(method))
		encoder, ok := encoders[method];
		if !ok {
			continue
		}

		return method, encoder
	}

	return "", nil
}