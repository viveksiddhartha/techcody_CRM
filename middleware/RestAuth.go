package middleware

import (
	"fmt"
	"log"
	"net/http"
	"svcrm/authenticate"
)

func GatedRestAuthHandler(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		secureCookieMap, err := authenticate.ReadSecureCookieValues(w, r)

		log.Printf("secure cookie map: %+v", secureCookieMap)

		if err != nil {
			log.Print(err)
		}

		//fmt.Printf("secure cookie contents: %+v\n", secureCookieMap)

		// Check if the sid key which is used to store the session id value
		// has been populated in the map using the comma ok idiom
		if _, ok := secureCookieMap["sid"]; ok == true {

			svSession, err := authenticate.SessionStore.Get(r, "sv_crm-session")

			fmt.Printf("SV_CRM session values: %+v\n", svSession.Values)
			if err != nil {
				log.Print(err)
				return
			}

			// Check if the session id stored in the secure cookie matches
			// the id and username on the server-side session
			if svSession.Values["sessionID"] == secureCookieMap["sid"] && svSession.Values["username"] == secureCookieMap["username"] {
				next(w, r)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("No ValidSession identified"))
			}

		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("No Session identified"))

		}

	})

}
