package main

import (
	"Vitor-Bichara/servidorGO/iogames"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	log.Print("Starting server...")

	err := iogames.CreateDBConnection()

	if err != nil {
		panic(err)
	}

	iogames.StartValidator()

	StartServer()

}

//StartServer .
func StartServer() {

	r := mux.NewRouter()

	iogames.SetupRoutes(r)
	iogames.SetupSessions()

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./iogames/public"))))
	http.ListenAndServe(iogames.SERVER_HOST, r)

}
