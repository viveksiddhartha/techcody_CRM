package handlers

import (
	"SV_CRM/common"
	"SV_CRM/common/datastore"
	"SV_CRM/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Users struct {
	models.User
}

//SignUpHandler for sinup hanlder
func SignUpHandler(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}
		fmt.Println("This is the reqBody %v", reqBody)

		u := models.User{}
		json.Unmarshal(reqBody, &u)
		fmt.Println("this is var u models.user %v", u.Email)

		//UserDetails := models.NewUser(r.FormValue(u.Username), r.FormValue(u.FirstName), r.FormValue(u.LastName), r.FormValue(u.Email), r.FormValue(u.PasswordHash))

		//Optional statement to experiment the new way

		err = datastore.CreateUser(&u)

		if err != nil {
			log.Print(err)
		}

	})
}

/*
	fmt.Println("End point hit")
	var user models.User
	json.NewEncoder(w).Encode(&user)



}
*/
