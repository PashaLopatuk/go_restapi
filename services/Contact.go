package services

import (
	"encoding/json"
	"example/go_restapi/models"
	"example/go_restapi/repository"
	"fmt"
	"net/http"
)

func GetContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(repository.Contacts)
}

func PostContact(w http.ResponseWriter, r *http.Request) {
	var contact models.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	repository.Contacts = append(repository.Contacts, contact)
	fmt.Fprint(w, "post new person '%v'", contact)
}
