package handlers

import (
	"net/http"
	"svcrm/authenticate"
)

func LogOutCRM(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)

}
