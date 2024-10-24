package fw

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

var fieldReplacer = strings.NewReplacer(
	"\n", "\\n",
	"\r", "\\r")

var dataReplacer = strings.NewReplacer(
	"\n", "\ndata:",
	"\r", "\\r")

type SSEEvent struct {
	Event string `json:"event"`
	Id    string `json:"id"`
	Retry uint64 `json:"retry"`
	Data  []byte `json:"data"`
}

func EncodeSSEEvent(w io.Writer, e *SSEEvent) error {
	if err := writeSSEId(w, e.Id); err != nil {
		return err
	}
	if err := writeSSEEvent(w, e.Event); err != nil {
		return err
	}
	if err := writeSSERetry(w, e.Retry); err != nil {
		return err
	}
	if err := writeSSEData(w, e.Data); err != nil {
		return err
	}
	return nil
}

func writeSSEId(w io.Writer, id string) error {
	if len(id) > 0 {
		if _, err := fmt.Fprintf(w, "id:%v\n", fieldReplacer.Replace(id)); err != nil {
			return err
		}
	}

	return nil
}

func writeSSEEvent(w io.Writer, event string) error {
	if len(event) > 0 {
		if _, err := fmt.Fprintf(w, "event:%v\n", fieldReplacer.Replace(event)); err != nil {
			return err
		}
	}

	return nil
}

func writeSSERetry(w io.Writer, retry uint64) error {
	if retry > 0 {
		if _, err := fmt.Fprintf(w, "retry:%v\n", strconv.FormatUint(retry, 10)); err != nil {
			return err
		}
	}

	return nil
}

func writeSSEData(w io.Writer, data []byte) error {
	if len(data) > 0 {
		if _, err := fmt.Fprintf(w, "data: %v\n\n", dataReplacer.Replace(string(data))); err != nil {
			return err
		}
	}

	return nil
}
