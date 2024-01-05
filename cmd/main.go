package main

import (
	"encoding/json"
	"example/go_restapi/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/contact", handler.ContactHandler)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode("Hello world")
	//fmt.Fprintf(w, "Привет!")
}
