package main

import (
	"SVCRM/common"
	"SVCRM/common/datastore"
	"SVCRM/handlers"
	"SVCRM/middleware"
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
	router.Handle("/api/crm/login", handlers.LoginEntity(&env))
	router.Handle("/api/crm/entity", handlers.EntityCreate(&env))
	router.HandleFunc("/api/crm/logout", handlers.LogOutCRM)

	//========Cookie based Authenticiation
	router.Handle("/api/crm/profile", middleware.GatedRestAuthHandler(handlers.ProfileCreate(&env)))
	router.Handle("/api/crm/profilelist", middleware.GatedRestAuthHandler(handlers.GetAllProfile(&env)))
	router.Handle("/api/crm/updateprofile", middleware.GatedRestAuthHandler(handlers.UpdateProfile(&env)))
	router.Handle("/api/crm/updateentity", middleware.GatedRestAuthHandler(handlers.UpdateEntity(&env)))
	router.Handle("/api/crm/getprofile", middleware.GatedRestAuthHandler(handlers.GetAllEntity(&env)))
	router.Handle("/api/crm/contract", middleware.GatedRestAuthHandler(handlers.CreateContract(&env)))

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
