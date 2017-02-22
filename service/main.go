package main

//-- Imports
import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//-- Constants
const port int = 9000

type Hello struct {
	Msg string `json:"msg"`
}

//-- Handlers
func HomeHandler(response http.ResponseWriter, request *http.Request) {

	msg := Hello{Msg: "Hello world!"}

	json.NewEncoder(response).Encode(msg)
}

//-- Entry point of the app
func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	//Pass mux to http
	http.Handle("/", r)
	portStr := strconv.Itoa(port)

	log.Print("Starting web service on " + portStr + " ... ")

	err := http.ListenAndServe(":"+portStr, nil)

	if err != nil {
		log.Fatal("Could not start the web service: ", err)
	}
}
