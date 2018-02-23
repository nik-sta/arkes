package main

import "encoding/json"
import "log"
import "net/http"
import "github.com/gorilla/mux"

type Greeting struct {
	Message string `json:"message,omitempty"`
}

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/hello", SayHello).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func SayHello(w http.ResponseWriter, r *http.Request)  {
	greeting := Greeting{"Hello :)"}
	json.NewEncoder(w).Encode(greeting)
}
