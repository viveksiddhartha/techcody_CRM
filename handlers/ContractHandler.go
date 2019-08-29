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

func CreateContract(e *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		svSession, err := authenticate.SessionStore.Get(r, "sv_crm-session")
		if err != nil {
			log.Print(err)
			return
		}
		username := svSession.Values["username"]
		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.Contract{}

		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal([]byte(reqBody), &u)
		fmt.Printf("%+v\n", u)

		if u.CoEntityID == username {
			_, error := datastore.GetEntityDetailsByCoEntityId(u.CoEntityID)
			if error == sql.ErrNoRows {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				ErrorResponse := map[string]string{
					"ErrorCode":   "NO_PROFILE_EXIST",
					"DESCRIPTION": "Not profile exist against the entity",
				}
				json.NewEncoder(w).Encode(ErrorResponse)
			} else {
				_, error := datastore.GetContractDetailsByCoEntityID(u.CoEntityID)
				if error == sql.ErrNoRows {
					//This code used to marshal the json
					var b []byte
					b, err = json.Marshal(&u)

					u.JsonObject = b

					datastore.CreateContract(&u)
					datastore.CreateAllocation(&u)

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(u)

					fmt.Println("Value of profile %v \n")

				} else {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					ErrorResponse := map[string]string{
						"ErrorCode":   "CONTRACT_ALREADY_EXIST",
						"DESCRIPTION": "Contract for this Profile already Exist",
					}
					json.NewEncoder(w).Encode(ErrorResponse)

				}
			}

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please login again"))
		}
	})
}
