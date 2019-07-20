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

type Profiles struct {
	models.Profile
}

//SignUpHandler for sinup hanlder
func ProfileCreate(e *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print("Encountered error when attempting to read the request body: ", err)
		}
		
		fmt.Println("This is the reqBody %v", reqBody)

		u := models.Profile{}
		json.Unmarshal(reqBody, &u)
		fmt.Println("this is var u models.Profile %v", u.Email)

		//ProfileDetails := models.NewProfile(r.FormValue(u.Profilename), r.FormValue(u.FirstName), r.FormValue(u.LastName), r.FormValue(u.Email), r.FormValue(u.PasswordHash))

		//Optional statement to experiment the new way

		err = datastore.CreateProfile(&u)

		if err != nil {
			log.Print(err)
		}

	})
}
