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

type CoEntity struct {
	models.CoEntity
}

func EntityCreate(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}
		fmt.Println("This is the reqBody %v", reqBody)

		u := models.CoEntity{}
		json.Unmarshal(reqBody, &u)
		fmt.Println("this is var u models.Profile %v", u)
		//ProfileDetails := models.NewProfile(r.FormValue(u.Profilename), r.FormValue(u.FirstName), r.FormValue(u.LastName), r.FormValue(u.Email), r.FormValue(u.PasswordHash))

		//Optional statement to experiment the new way

		err = datastore.EntityCreate(&u)
		if err != nil {
			log.Print(err)
		}

	})
}
