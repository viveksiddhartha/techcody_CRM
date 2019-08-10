package handlers

import (
	"SV_CRM/authenticate"
	"SV_CRM/common"
	"SV_CRM/common/datastore"
	"SV_CRM/common/utility"
	"SV_CRM/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LoginSts struct {
	models.LoginSt
	models.SessionSt
	datastore.Datastore
	Error error
}

func LoginEntity(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		ul := models.LoginSt{}
		if err != nil {
			log.Print(err)
		}

		json.Unmarshal(reqBody, &ul)

		fmt.Printf("value is %v \n", ul)

		if ul.Username == "" {
			fmt.Printf("\n User name is missing \n")
			w.Write([]byte("500 - User name is missing!"))
		}
		if ul.Password == "" {
			fmt.Printf("Password name is missing \n")
			w.Write([]byte("500 - User name is missing!"))
		} else {
			fmt.Println("Check 1 Ready for Authentication")
			authResult := authenticate.VerifyCredentials(env, ul.Username, ul.Password)
			fmt.Println("Check 2")
			fmt.Println("auth result: ", authResult)

			if authResult == true {

				sessionID := utility.GenerateUUID()
				fmt.Println("sessid: ", sessionID)

				err = authenticate.CreateSecureCookie(&ul, sessionID, w, r)
				if err != nil {
					log.Print("Encountered error when attempting to create secure cookie: ", err)
					//http.Redirect(w, r, "/login", 302)
					return

				}
				fmt.Println("Check 2 Ready for Authentication")

				err = authenticate.CreateUserSession(ul.Username, sessionID, w, r)
				if err != nil {
					log.Print("Encountered error when attempting to create user session: ", err)
					w.WriteHeader(http.StatusInternalServerError)
					w.Write([]byte("500 - login failed"))
					return

				}

				usession := map[string]string{
					"Username": ul.Username,
					"Session":  sessionID,
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				//w.Write([]byte("200 - authenticated successfully"))
				json.NewEncoder(w).Encode(usession)

				return

			} else {

				w.Write([]byte("500 - authenticated failed"))
				return

			}

		}
	})
}
