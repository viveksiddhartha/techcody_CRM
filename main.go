package main

import (
	"SV_CRM/common/datastore"
	"SV_CRM/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (

	//WEBSERVERPORT port for starting the server side
	WEBSERVERPORT = ":8080"
)

func main() {

	//Connection for Database
	datastore.DBConn()

	fmt.Println("Connected successfully")

	//defer db.Close()

	//New router created for handler function
	router := mux.NewRouter()

	//New handlers for CRM profile handling
	http.Handle("/", router)
	//Home handler will provide the welcome message on index page
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/signup", handlers.SignUpHandler)

	//Lister defined for end point
	/*
		loggedRouter := ghandlers.LoggingHandler(os.Stdout, router)
		stdChain := alice.New(middleware.PanicRecoveryHandler)
		http.Handle("/", stdChain.Then(loggedRouter))
	*/
	http.ListenAndServe(WEBSERVERPORT, nil)

}
