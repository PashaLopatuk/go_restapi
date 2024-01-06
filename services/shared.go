package services

import (
	"net/http"
)

func ErrorNotAllowed(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode("Method not Allowed")
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("\"Method not Allowed\""))
}

func ErrorInternalServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("\"Internal Server error\""))
}
