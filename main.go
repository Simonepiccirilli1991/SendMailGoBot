package sendmailgobot

import (
	"SendMailGoBot/service"
	"fmt"
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

func handler(w http.ResponseWriter, r *http.Request) {

	checkPeding = service.GetMailPending()
	fmt.Fprintf(w, "Hello World!")
}
