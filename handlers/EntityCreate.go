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

type CoEntity struct {
	models.CoEntity
	datastore.Datastore
	datastore.RDatastore
	Error error
}

func EntityCreate(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.CoEntity{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)

		_, error := datastore.GetEntityDetailsByCoEntityId(u.CoEntityId)
		//To enable the redis database get query
		//values, error := datastore.GetUserRedis(u.CoEntityId)

		if error == sql.ErrNoRows {

			//ProfileDetails := models.NewProfile(r.FormValue(u.Profilename), r.FormValue(u.FirstName), r.FormValue(u.LastName), r.FormValue(u.Email), r.FormValue(u.PasswordHash))

			err = datastore.EntityCreate(&u)
			if err != nil {
				log.Print(err)
			}
			EntityValue := map[string]string{
				"UserName":    u.CoEntityId,
				"CompanyName": u.CompanyNm,
				"Status":      "Draft",
			}

			//Optional statement to experiment the new way... To enable the entity creation in Redis database
			/* 			err = datastore.CreateEntityRedis(&u)
			   			if err != nil {
			   				log.Print(err)
			   			}
			*/
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			//w.Write([]byte("200 - authenticated successfully"))
			json.NewEncoder(w).Encode(EntityValue)

		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Printf("User already exist %v \n", u.CoEntityId)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - User already exist!"))

		}

	})
}

/*
func EntityUpdate(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.Profile{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)

		//db.QueryRow("select * FROM coentity WHERE Status in ('0','1') and CoEntityID = $1", u.CoEntityId)

		_, error := datastore.GetProfileDetailsByProfilename(u.Profilename)
		if error == sql.ErrNoRows {

			fmt.Printf("No User exist %v \n", u.Profilename)
			w.Write([]byte("200 - No record exist!"))

		} else {
			err = datastore.EntityCreate(&u)
			if err != nil {
				log.Print(err)
			}
			w.Write([]byte("200 - Profile has been Created successfully!"))

		}

	})
}
*/
