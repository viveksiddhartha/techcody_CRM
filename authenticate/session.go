package authenticate

import (
	"SVCRM/common/datastore"
	"SVCRM/models"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type LoginSts struct {
	models.CoEntity
	models.LoginSt
	datastore.Datastore
	Error error
}

var SessionStore = sessions.NewCookieStore([]byte("SV_CRM_SECRET_HASHKEY_S"))

//var SessionStore *sessions.FilesystemStore

func CreateUserSession(username, sessionID string, w http.ResponseWriter, r *http.Request) error {

	svSession, err := SessionStore.Get(r, "sv_crm-session")
	ue, err := datastore.GetEntityDetailsByCoEntityId(username)
	if err != nil {
		log.Print("Encountered error when attempting to fetch user record: ", err)
		//http.Redirect(w, r, "/login", 302)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Record not exist!"))
	}
	if err != nil {
		log.Print(err)
	}

	svSession.Values["sessionID"] = sessionID
	svSession.Values["username"] = ue.CoEntityId
	svSession.Values["CompanyNm"] = ue.CompanyNm
	svSession.Values["AliasNm"] = ue.AliasNm
	svSession.Values["email"] = ue.Email
	svSession.Values["uuid"] = ue.UUID

	err = svSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	svSession, err := SessionStore.Get(r, "sv_crm-session")

	if err != nil {
		log.Print(err)
	}

	svSession.Options.MaxAge = -1
	svSession.Save(r, w)
}

/*
func init() {
	if _, err := os.Stat("/tmp/gopherface-sessions"); os.IsNotExist(err) {
		os.Mkdir("/tmp/gopherface-sessions", 711)
	}
	SessionStore = sessions.NewFilesystemStore("/tmp/gopherface-sessions")
}
*/
