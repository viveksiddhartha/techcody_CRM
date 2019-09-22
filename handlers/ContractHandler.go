package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"svcrm/authenticate"
	"svcrm/common"
	"svcrm/common/datastore"
	"svcrm/models"
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
				_, error := datastore.GetContractsByCoEntityID(u.CoEntityID)
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

func GetAllContracts(e *common.Env) http.HandlerFunc {
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
					"ErrorCode":   "NO_ENTITY_EXIST",
					"DESCRIPTION": "No Entity exist for given CoEntityID",
				}
				json.NewEncoder(w).Encode(ErrorResponse)
			} else {
				_, error := datastore.GetContractsByCoEntityID(u.CoEntityID)
				if error == sql.ErrNoRows {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					ErrorResponse := map[string]string{
						"ErrorCode":   "NO_CONTRACT_EXIST",
						"DESCRIPTION": "No Contract exist for given CoEntityID",
					}
					json.NewEncoder(w).Encode(ErrorResponse)
				} else {
					up, err := datastore.GetAllContractByCoEntity(u.CoEntityID)
					if err != nil {
						log.Print(err)
					}

					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					json.NewEncoder(w).Encode(up)

					fmt.Println("Value of profile %v \n")

				}
			}

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please login again"))
		}
	})
}

func GetAllocationByContractID(e *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		svSession, err := authenticate.SessionStore.Get(r, "sv_crm-session")
		if err != nil {
			log.Print(err)
			return
		}
		username := svSession.Values["username"]
		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.AllocationList{}

		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal([]byte(reqBody), &u)
		fmt.Printf("%+v\n", u)

		if u.CoEntityID == username {
			_, error := datastore.GetContractsByCoEntityID(u.CoEntityID)
			if error == sql.ErrNoRows {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				ErrorResponse := map[string]string{
					"ErrorCode":   "NO_CONTRACT_EXIST",
					"DESCRIPTION": "No Contract exist for given CoEntityID",
				}
				json.NewEncoder(w).Encode(ErrorResponse)
			} else {
				_, error := datastore.GetContractsByCoEntityID(u.CoEntityID)
				if error == sql.ErrNoRows {
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					ErrorResponse := map[string]string{
						"ErrorCode":   "NO_CONTRACT_EXIST",
						"DESCRIPTION": "No Contract exist for given CoEntityID",
					}
					json.NewEncoder(w).Encode(ErrorResponse)
				} else {
					AllocationList, err := datastore.GetAllocationByCoEntityContractID(u.CoEntityID)
					if err != nil {
						log.Print("Encountered error when Fetching data from DB ", err)
					}
					w.Header().Set("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					//w.Write([]byte("200 - authenticated successfully"))
					json.NewEncoder(w).Encode(AllocationList)
				}
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please login again"))
		}
	})
}
