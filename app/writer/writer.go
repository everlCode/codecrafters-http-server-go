package writer

import (
	"net"

	"github.com/codecrafters-io/http-server-starter-go/app/parser"
)

type Writer struct {
	Conn net.Conn
}

func New(conn net.Conn) *Writer {
	return &Writer{
		Conn: conn,
	}
}

func (w *Writer) Write(r *parser.HttpResponse) {
	w.Conn.Write(r.Marshal())
}