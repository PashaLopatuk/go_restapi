package repository

import "example/go_restapi/models"

var Contacts []models.Contact

func ReadAllContactsRepository() []models.Contact {
	return Contacts
}

func ReadOneContactRepository(id int) models.Contact {
	return Contacts[id]
}

func CreateOneContactsRepository(contact models.Contact) {
	Contacts = append(Contacts, contact)
}
