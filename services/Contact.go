package services

import (
	"encoding/json"
	"example/go_restapi/models"
	"example/go_restapi/repository"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all")
	json.NewEncoder(w).Encode(repository.ReadAllContactsRepository())
}

func GetOneContacts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//log.Fatal("Id should be a valid integer!")
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	json.NewEncoder(w).Encode(repository.ReadOneContactRepository(id))
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create")
	var contact models.ContactReq
	err := json.NewDecoder(r.Body).Decode(&contact)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	repository.CreateContactsRepository(contact)
	//repository.Contacts = append(repository.Contacts, contact)
	fmt.Fprint(w, "Created new contact '%v'", contact)
}

func UpdateContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update")
	var contact models.ContactReq
	vars := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&contact)
	id, errId := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else if errId != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}

	repository.UpdateContactRepository(contact, id)
	//repository.Contacts = append(repository.Contacts, contact)
	fmt.Fprint(w, "Updated contact '%v'", contact)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete")
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		//log.Fatal("Id should be a valid integer!")
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
	}
	repository.DeleteContactRepository(id)
}

//
//func GetContactsOne(w http.ResponseWriter, r *http.Request) {
//	//json.NewEncoder(w).Encode(repository.Contacts[r.Body])
//	fmt.Println(r.Body)
//}
//
//func GetContacts(w http.ResponseWriter, r *http.Request) {
//	id := r.URL.Query().Get("id")
//	fmt.Println("id => ", id)
//	json.NewEncoder(w).Encode(repository.Contacts)
//}
//
//func PostContact(w http.ResponseWriter, r *http.Request) {
//	var contact models.Contact
//	err := json.NewDecoder(r.Body).Decode(&contact)
//
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//
//	repository.Contacts = append(repository.Contacts, contact)
//	fmt.Fprint(w, "post new person '%v'", contact)
//}
