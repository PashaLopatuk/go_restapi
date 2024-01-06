package main

import (
	"encoding/json"
	"example/go_restapi/handler"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//type Handler struct{}
//
//func (h *Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	fmt.Println("request: ", req.RequestURI)
//	_, err := w.Write([]byte("Hello world"))
//	if err != nil {
//		log.Println(err)
//	}
//}

func main() {
	//log.Println(http.ListenAndServe(":8080", &Handler{}))
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc(
		"/contacts",
		handler.ContactHandler(handler.GETALL),
	).Methods(http.MethodGet)
	router.HandleFunc(
		"/contacts/{id:[0-9]+}",
		handler.ContactHandler(handler.GETONE),
	).Methods(http.MethodGet)
	router.HandleFunc(
		"/contacts",
		handler.ContactHandler(handler.CREATE),
	).Methods(http.MethodPost)
	router.HandleFunc(
		"/contacts/{id:[0-9]+}",
		handler.ContactHandler(handler.UPDATE),
	).Methods(http.MethodPatch)
	router.HandleFunc(
		"/contacts/{id:[0-9]+}",
		handler.ContactHandler(handler.DELETE),
	).Methods(http.MethodDelete)

	err := http.ListenAndServe("localhost:8080", router)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("Server started!")

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode("Hello world")
	fmt.Println(r.URL.Query().Get("id"))

	//fmt.Fprintf(w, "Hello world!")
}
