package authenticate

import (
	"SVCRM/models"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
)

/*
type LoginStsC struct {
	models.LoginSt
	datastore.Datastore
	Error error
}
*/
var hashSKey = []byte("SV_CRM_SECRET_HASHKEY_S")
var blockKey = []byte("SV_CRM_SECRET_BLOCKKEY_S")
var s = securecookie.New(hashSKey, blockKey)

func CreateSecureCookie(ul *models.LoginSt, sessionID string, w http.ResponseWriter, r *http.Request) error {

	value := map[string]string{
		"username": ul.Username,
		"sid":      sessionID,
	}

	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)

	} else {
		return err
	}
	return nil

}

func ReadSecureCookieValues(w http.ResponseWriter, r *http.Request) (map[string]string, error) {
	if cookie, err := r.Cookie("session"); err == nil {
		value := make(map[string]string)
		if err = s.Decode("session", cookie.Value, &value); err == nil {
			return value, nil
		} else {
			return nil, err
		}
	} else {
		return nil, nil
	}
}

func ExpireSecureCookie(w http.ResponseWriter, r *http.Request) {

	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	w.Header().Set("Cache-Control", "no-cache, private, max-age=0")
	w.Header().Set("Expires", time.Unix(0, 0).Format(http.TimeFormat))
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("X-Accel-Expires", "0")

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("logout successfully"))

}
