package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>'Ello from Go!</h1>")
}
