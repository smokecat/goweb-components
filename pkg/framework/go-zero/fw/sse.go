package fw

import (
	"io"
	"net/http"
)

type sseWriter interface {
	io.Writer
	http.Flusher
}

type SSEStream struct {
	writer sseWriter
	// writer http.ResponseWriter
}

func NewSSEStream(w http.ResponseWriter, r *http.Request) *SSEStream {
	w.Header().Set("Content-Type", "text/event-stream")
	writer, _ := w.(sseWriter)
	return &SSEStream{
		writer: writer,
	}
}

func (s *SSEStream) Publish(event *SSEEvent) error {
	// fmt.Fprintf(s.writer, "data: %s\n\n", event.Data)
	// return nil
	err := EncodeSSEEvent(s.writer, event)
	if err != nil {
		return err
	}
	s.writer.(http.Flusher).Flush()
	return nil
}
