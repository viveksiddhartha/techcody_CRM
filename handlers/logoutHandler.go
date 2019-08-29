package handlers

import (
	"SV_CRM/authenticate"
	"net/http"
)

func LogOutCRM(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)

}
