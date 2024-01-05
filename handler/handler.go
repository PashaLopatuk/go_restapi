package handler

import (
	"example/go_restapi/services"
	"net/http"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		services.GetContacts(w, r)
	case http.MethodPost:
		services.PostContact(w, r)
	default:
		services.ErrorNotAllowed(w, r)
	}
}
