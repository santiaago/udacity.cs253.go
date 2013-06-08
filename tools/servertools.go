package tools

import (
	"io"
	"net/http"
)


func Error404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "404: Not Found")
}
