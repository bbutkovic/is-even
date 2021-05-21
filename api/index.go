package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	number, err := strconv.Atoi(r.URL.Query().Get("number"))

	if err != nil {
		fmt.Fprintf(w, "number not integer")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if number%2 == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "true")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "false")
	}
}
