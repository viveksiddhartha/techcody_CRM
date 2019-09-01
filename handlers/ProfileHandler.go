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

type Profiles struct {
	models.Profile
	datastore.Datastore
	Error error
}

//SignUpHandler for sinup hanlder
func ProfileCreate(e *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		svSession, err := authenticate.SessionStore.Get(r, "sv_crm-session")
		if err != nil {
			log.Print(err)
			return
		}
		username := svSession.Values["username"]

		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.Profile{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)

		if u.CoEntityID == username {

			_, error := datastore.GetProfileDetailsByCoEntityProfilename(u.CoEntityID, u.Profilename)
			if error == sql.ErrNoRows {

				err = datastore.CreateProfile(&u)

				if err != nil {
					log.Print(err)
				}
				UProfile := map[string]string{
					"Profile":   u.Profilename,
					"FirstName": u.FirstName,
					"LastName":  u.LastName,
					"Email":     u.Email,
					"Phone":     u.ContactNo,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				//w.Write([]byte("200 - authenticated successfully"))
				json.NewEncoder(w).Encode(UProfile)

			} else {
				fmt.Printf("Profile already exist %v \n", u.Profilename)

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("500 - User already exist!"))
			}
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please login again"))
		}
	})
}

//Update profile for sinup hanlder
func UpdateProfile(e *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		svSession, err := authenticate.SessionStore.Get(r, "sv_crm-session")
		if err != nil {
			log.Print(err)
			return
		}
		username := svSession.Values["username"]

		fmt.Println("User name from Cookie%s\t", username)

		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.Profile{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)

		if u.CoEntityID == username {
			_, error := datastore.GetProfileDetailsByCoEntityProfilename(u.CoEntityID, u.Profilename)
			if error == sql.ErrNoRows {

				fmt.Printf("Profile does not exist %v \n", u.Profilename)

				w.Write([]byte("200 - No user exist!"))

			} else {
				err = datastore.UpdateProfileByProfileID(&u)
				if err != nil {
					log.Print(err)
				}
				UProfile := map[string]string{
					"Profile":   u.Profilename,
					"FirstName": u.FirstName,
					"LastName":  u.LastName,
					"Email":     u.Email,
					"Phone":     u.ContactNo,
				}
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				//w.Write([]byte("200 - authenticated successfully"))
				json.NewEncoder(w).Encode(UProfile)
			}

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Please login again"))
		}

	})
}

func GetAllProfile(e *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		svSession, err := authenticate.SessionStore.Get(r, "sv_crm-session")
		if err != nil {
			log.Print(err)
			return
		}
		username := svSession.Values["username"]

		reqBody, err := ioutil.ReadAll(r.Body)
		u := models.Profile{}
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}

		json.Unmarshal(reqBody, &u)

		if u.CoEntityID == username {

			_, error := datastore.GetProfileDetailsByCoEntity(u.CoEntityID)
			if error == sql.ErrNoRows {
				fmt.Printf("Profile already exist %v \n", u.Profilename)

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
				up, err := datastore.GetAllProfileDetailsByCoEntity(u.CoEntityID)
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
