package authenticate

import (
	"SV_CRM/common/datastore"
	"SV_CRM/models"
	"fmt"
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

var SessionStore = sessions.NewCookieStore([]byte("something-very-secret"))

//var SessionStore *sessions.FilesystemStore

func CreateUserSession(username, sessionID string, w http.ResponseWriter, r *http.Request) error {

	fmt.Println("Check 1 Ready for Session")
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
	fmt.Println("Check 2 Ready for Session %v \n", ue)

	svSession.Values["sessionID"] = sessionID
	svSession.Values["username"] = ue.CoEntityId
	svSession.Values["CompanyNm"] = ue.CompanyNm
	svSession.Values["AliasNm"] = ue.AliasNm
	svSession.Values["email"] = ue.Email
	svSession.Values["uuid"] = ue.UUID

	fmt.Println("Check 3 Ready for Session %v \n\n\n\n\n", svSession)
	fmt.Println("Check 4 Ready for Session %v \n\n\n\n\n", r)
	fmt.Println("Check 5 Ready for Session %v \n\n\n\n\n", w)

	err = svSession.Save(r, w)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func ExpireUserSession(w http.ResponseWriter, r *http.Request) {
	gfSession, err := SessionStore.Get(r, "sv_crm-session")

	if err != nil {
		log.Print(err)
	}

	gfSession.Options.MaxAge = -1
	gfSession.Save(r, w)
}

/*
func init() {
	if _, err := os.Stat("/tmp/sv_crm-sessions"); os.IsNotExist(err) {
		os.Mkdir("/tmp/sv_crm-sessions", 711)
	}
	SessionStore = sessions.NewFilesystemStore("/tmp/sv_crm-sessions", []byte(os.Getenv("SV_CRM_SECRET_HASHKEY_S")))
}
*/
