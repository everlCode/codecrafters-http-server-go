package httpx

type Header struct {
	Name  string
	Value string
}

const (
	HeaderUserAgent = "User-Agent"
	ContentLength   = "Content-Length"
	ContentType   = "Content-Type"
)
