package authenticate

import (
	"SVCRM/common"
	"SVCRM/common/datastore"
	"SVCRM/common/utility"
	"SVCRM/models"
	"fmt"
	"log"
	"strings"
)

//CoEntity value story
type CoEntity struct {
	models.LoginSt
	datastore.Datastore
	Error error
}

//VerifyCredentials function
func VerifyCredentials(e *common.Env, username string, password string) bool {

	u, err := datastore.GetEntityDetailsByCoEntityIdForPassword(username)
	if u == nil {
		return false
	}

	if err != nil {
		log.Print(err)
	}
	pwh := utility.SHA256OfString(password)

	fmt.Println("auth result: ", u.CoEntityId)
	fmt.Println("auth result: ", username)
	fmt.Println("auth result: ", u.Password)
	fmt.Println("auth result: ", pwh)

	fmt.Println("auth result: ", username)
	fmt.Println("auth result: ", u.CoEntityId)
	fmt.Println("auth result: ", pwh)
	fmt.Println("auth result: ", u.Password)

	if strings.ToLower(username) == strings.ToLower(u.CoEntityId) && utility.SHA256OfString(password) == u.Password {
		log.Println("Successful login attempt from user: ", u.CoEntityId)
		return true
	} else {
		log.Println("Unsuccessful login attempt from user: ", u.CoEntityId)
		return false
	}
}
