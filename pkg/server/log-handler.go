package server

import (
	"log"
	"net/http"
)

func NewLoggingHandler(next http.Handler) *loggingHandler {
	return &loggingHandler{next: next}
}

type loggingHandler struct {
	next http.Handler
}

func (h *loggingHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	log.Printf("%s %s %s", request.RemoteAddr, request.Method, request.URL.Path)
	h.next.ServeHTTP(writer, request)
}

var _ http.Handler = (*loggingHandler)(nil)
