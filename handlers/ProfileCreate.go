package handlers

import (
	"SV_CRM/common"
	"SV_CRM/common/datastore"
	"SV_CRM/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Profiles struct {
	models.Profile
	datastore.Datastore
	Error error
}

//SignUpHandler for sinup hanlder
func ProfileCreate(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.Profile{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)

		_, error := datastore.GetProfileDetailsByProfilCoEntity(u.Profilename, u.CoEntityID)
		if error == sql.ErrNoRows {

			err = datastore.CreateProfile(&u)

			if err != nil {
				log.Print(err)
			}
		} else {
			fmt.Printf("Profile already exist %v \n", u.Profilename)

			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - User already exist!"))
		}

	})
}

//Update profile for sinup hanlder
func UpdateProfile(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.Profile{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)

		_, error := datastore.GetProfileDetailsByProfilename(u.Profilename)
		if error == sql.ErrNoRows {

			fmt.Printf("Profile does not exist %v \n", u.Profilename)

			w.Write([]byte("200 - No user exist!"))

		} else {
			err = datastore.UpdateProfileByProfileID(&u)

			if err != nil {
				log.Print(err)
			}

		}

	})
}
