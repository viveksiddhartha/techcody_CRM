package authenticate

import (
	"SV_CRM/models"
	"fmt"
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
	fmt.Printf("Check 3 Ready for Create Secure Cookie %v %v \n ", sessionID, ul.Username)
	value := map[string]string{
		"username": ul.Username,
		"sid":      sessionID,
	}
	fmt.Println("Check 4 Ready for Create Secure Cookie %v \n", value)

	if encoded, err := s.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:     "session",
			Value:    encoded,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
		}
		fmt.Println("Check 6 Ready for SetCookie %v %v \n", cookie, w)
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
	http.Redirect(w, r, "/login", 301)
}
