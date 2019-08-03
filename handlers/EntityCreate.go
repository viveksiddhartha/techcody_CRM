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

		//db.QueryRow("select * FROM coentity WHERE Status in ('0','1') and CoEntityID = $1", u.CoEntityId)

		_, error := datastore.GetEntityDetailsByCoEntityId(u.CoEntityId)
		if error == sql.ErrNoRows {

			//ProfileDetails := models.NewProfile(r.FormValue(u.Profilename), r.FormValue(u.FirstName), r.FormValue(u.LastName), r.FormValue(u.Email), r.FormValue(u.PasswordHash))

			//Optional statement to experiment the new way

			err = datastore.EntityCreate(&u)
			if err != nil {
				log.Print(err)
			}
			w.Write([]byte("200 - Profile has been Created successfully!"))

		} else {
			fmt.Printf("User already exist %v \n", u.CoEntityId)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - User already exist!"))

		}

	})
}
