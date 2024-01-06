package handler

import (
	"example/go_restapi/services"
	"net/http"
)

func ContactHandler(functionType string) func(w http.ResponseWriter, r *http.Request) {
	switch functionType {
	case GETONE:
		return services.GetOneContacts
	case GETALL:
		return services.GetAllContacts
	case CREATE:
		return services.CreateContact
	case UPDATE:
		return services.UpdateContact
	case DELETE:
		return services.DeleteContact
	}
	return nil
}

// GetContactHandler

// CreateContactHandler

// UpdateContactHandler

// DeleteContactHandler
