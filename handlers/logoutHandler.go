package handlers

import (
	"SVCRM/authenticate"
	"net/http"
)

func LogOutCRM(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)

}
