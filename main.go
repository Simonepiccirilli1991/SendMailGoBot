package main

import (
	"SendMailGoBot/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// The new router function creates the router and
// returns it to us. We can now use this function
// to instantiate and test the router outside of the main function
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/notify", handler).Methods("POST")

	return r
}
func main() {

	r := newRouter()
	http.ListenAndServe(":8080", r)

	if r != nil {
		log.Fatal("Se rotto")
	}

}

func handler(w http.ResponseWriter, r *http.Request) {

	mailDto, err := service.GetMailPending()
	if err != nil {
		log.Printf("no mail pending")
		fmt.Fprintf(w, "failed to get mail")
	}

	service.SendMail(*mailDto)

	log.Printf("Mail sended: %+v", mailDto)
	fmt.Fprintf(w, "mail sended succefully!")
}
