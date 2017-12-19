package httputil

import (
	"log"
	"net/http"
)

func WriteResponse(w http.ResponseWriter, b []byte) {
	_, err := w.Write(b)
	if err != nil {
		InternalServerError(w, err)
	}
	log.Printf("http status: %v", http.StatusText(http.StatusOK))
}

func InternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	statusText := http.StatusText(http.StatusInternalServerError)
	log.Printf("http status %v: %v", statusText, err)
}

func NotImplemented(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotImplemented)
	statusText := http.StatusText(http.StatusNotImplemented)
	log.Printf("http status: %v", statusText)
}
