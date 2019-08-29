package handlers

import (
	"SV_CRM/authenticate"
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
			fmt.Printf("Valid session does not exist %v \n", u.CoEntityId)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Valid session does not exist"))

		}

	})
}

func UpdateEntity(e *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		svSession, err := authenticate.SessionStore.Get(r, "sv_crm-session")
		if err != nil {
			log.Print(err)
			return
		}
		username := svSession.Values["username"]

		fmt.Println("\nUser name from Cookie%s\t\n", username)

		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.CoEntity{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)
		fmt.Println("\n Username from Request%s\t\n", u.CoEntityId)

		//db.QueryRow("select * FROM coentity WHERE Status in ('0','1') and CoEntityID = $1", u.CoEntityId)

		if u.CoEntityId == username {
			_, error := datastore.GetEntityDetailsByCoEntityId(u.CoEntityId)
			if error == sql.ErrNoRows {

				fmt.Printf("No User exist %v \n", u.CoEntityId)
				w.Write([]byte("200 - No record exist!"))

			} else {
				err = datastore.UpdateEntityByEntityID(&u)
				if err != nil {
					log.Print(err)
				}
				UEntity, err := datastore.GetEntityDetailsByCoEntityId(u.CoEntityId)
				if err != nil {
					log.Print(err)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				//w.Write([]byte("200 - authenticated successfully"))
				json.NewEncoder(w).Encode(UEntity)

			}
		} else {
			w.Header().Set("Content-Type", "application/json")
			fmt.Printf("Valid session does not exist \n")
			ErrorResponse := map[string]string{
				"ErrorCode":   "INVALID_SESSION",
				"DESCRIPTION": "Valid session does not exist",
			}
			json.NewEncoder(w).Encode(ErrorResponse)
		}

	})
}

func GetAllEntity(e *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		svSession, err := authenticate.SessionStore.Get(r, "sv_crm-session")
		if err != nil {
			log.Print(err)
			return
		}
		username := svSession.Values["username"]

		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.CoEntity{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)

		if u.CoEntityId == username {

			_, error := datastore.GetEntityDetailsByCoEntityId(u.CoEntityId)
			if error == sql.ErrNoRows {
				fmt.Printf("Profile already exist %v \n", u.CoEntityId)

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				ErrorResponse := map[string]string{
					"ErrorCode":   "NO_PROFILE_EXIST",
					"DESCRIPTION": "Not profile exist against the entity",
				}
				json.NewEncoder(w).Encode(ErrorResponse)

			} else if error != nil {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)

				json.NewEncoder(w).Encode(error)
				//w.Write([]byte("200 - authenticated successfully"))

			} else {
				up, err := datastore.GetEntityDetailsByCoEntityId(u.CoEntityId)
				if err != nil {
					log.Print(err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				//w.Write([]byte("200 - authenticated successfully"))
				json.NewEncoder(w).Encode(up)
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please login again"))
		}
	})
}
