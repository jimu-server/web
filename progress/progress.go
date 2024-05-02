package progress

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type Option func(progress *Progress)

type Progress struct {
	Status  int     `json:"status"`
	Mgs     string  `json:"msg"`
	Percent float64 `json:"percent"`
}

type WebProgress struct {
	w       http.ResponseWriter
	flusher http.Flusher
	percent float64
}

func NewProgress(w http.ResponseWriter) (*WebProgress, error) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	flusher, ok := w.(http.Flusher)
	if !ok {
		return nil, errors.New("not support")
	}
	progress := &WebProgress{
		w:       w,
		flusher: flusher,
	}
	return progress, nil
}

func (receiver *WebProgress) Progress(percent float64, message string, options ...Option) error {
	progress := Progress{
		Status:  http.StatusOK,
		Mgs:     message,
		Percent: percent * 0.01,
	}
	for _, o := range options {
		o(&progress)
	}
	return receiver.progress(progress)
}

func (receiver *WebProgress) progress(progress Progress) error {
	marshal, _ := jsoniter.Marshal(progress)
	value := string(marshal) + "\n"
	if _, err := receiver.w.Write([]byte(value)); err != nil {
		return err
	}
	receiver.flusher.Flush()
	return nil
}

func Error() Option {
	return func(progress *Progress) {
		progress.Status = http.StatusInternalServerError
	}
}
