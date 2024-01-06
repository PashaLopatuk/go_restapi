package repository

import (
	"example/go_restapi/models"
	"slices"
)

var Contacts []models.Contact

func ReadAllContactsRepository() []models.Contact {
	return Contacts
}

func ReadOneContactRepository(id int) models.Contact {
	return Contacts[id]
}

func CreateContactsRepository(contact models.ContactReq) {
	Contacts = append(Contacts, models.Contact{
		Id:          len(Contacts),
		Name:        contact.Name,
		Email:       contact.Email,
		PhoneNumber: contact.PhoneNumber,
	})
}

func UpdateContactRepository(contact models.ContactReq, id int) {
	idx := slices.IndexFunc(Contacts, func(c models.Contact) bool { return c.Id == id })
	Contacts[idx] = models.Contact{
		Id:          Contacts[idx].Id,
		Name:        contact.Name,
		Email:       contact.Email,
		PhoneNumber: contact.PhoneNumber,
	}
}

func DeleteContactRepository(id int) {
	idx := slices.IndexFunc(Contacts, func(c models.Contact) bool { return c.Id == id })
	Contacts = append(Contacts[:idx], Contacts[idx+1:]...)
	//Contacts[id] = nil
}
