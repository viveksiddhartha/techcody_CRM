package main

import (
	"SV_CRM/common"
	"SV_CRM/common/datastore"
	"SV_CRM/handlers"
	"SV_CRM/middleware"
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

	datastore.NewRedisDatastore()

	fmt.Println("Connected successfully")

	env := common.Env{}

	db := datastore.DBConn()
	defer db.Close()
	//New router created for handler function

	router := mux.NewRouter()

	//New handlers for CRM profile handling
	http.Handle("/", router)
	//Home handler will provide the welcome message on index page
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.Handle("/login", handlers.LoginEntity(&env))
	router.Handle("/entity", handlers.EntityCreate(&env))
	router.HandleFunc("/logout", handlers.LogOutCRM)

	//========Cookie based Authenticiation
	router.Handle("/profile", middleware.GatedRestAuthHandler(handlers.ProfileCreate(&env)))
	router.Handle("/profilelist", middleware.GatedRestAuthHandler(handlers.GetAllProfile(&env)))
	router.Handle("/updateprofile", middleware.GatedRestAuthHandler(handlers.UpdateProfile(&env)))
	router.Handle("/updateentity", middleware.GatedRestAuthHandler(handlers.UpdateEntity(&env)))
	router.Handle("/getprofile", middleware.GatedRestAuthHandler(handlers.GetAllEntity(&env)))
	router.Handle("/contract", middleware.GatedRestAuthHandler(handlers.CreateContract(&env)))

	//========Cookie based Authenticiation
	router.HandleFunc("/healthcheck", handlers.HealthCheckHandler)

	//router.Handle("/updateprofile", logger.Logger(handlers.UpdateProfile(&env)))

	//Lister defined for end point
	/*
		loggedRouter := ghandlers.LoggingHandler(os.Stdout, router)
		stdChain := alice.New(middleware.PanicRecoveryHandler)
		http.Handle("/", stdChain.Then(loggedRouter))
	*/
	http.ListenAndServe(WEBSERVERPORT, nil)

}
