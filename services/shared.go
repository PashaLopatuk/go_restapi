package services

import (
	"net/http"
)

func ErrorNotAllowed(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode("Method not Allowed")
}
