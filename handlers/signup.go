package handlers

import (
	"SV_CRM/common/datastore"
	"SV_CRM/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//SignUpForm for page struct

//SignUpHandler for sinup hanlder
func SignUpHandler(w http.ResponseWriter, r *http.Request) {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Print("Encountered error when attempting to read the request body: ", err)
	}
	fmt.Println("This is the reqBody %v", reqBody)

	var u models.User
	json.Unmarshal(reqBody, &u)
	fmt.Println("this is var u models.user %v", u)

	UserDetails := models.NewUser(r.FormValue("username"), r.FormValue("firstName"), r.FormValue("lastName"), r.FormValue("email"), r.FormValue("password"))

	fmt.Println("this is userdetails %v", u)
	err = datastore.CreateUser(UserDetails)

	if err != nil {
		log.Print(err)
	}

}

/*
	fmt.Println("End point hit")
	var user models.User
	json.NewEncoder(w).Encode(&user)



}
*/
