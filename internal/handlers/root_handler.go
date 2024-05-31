package handlers

import (
	"net/http"
)

func FaviconHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
