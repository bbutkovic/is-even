package handler

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	fmt.Fprintf(w, "%s", number)
}
