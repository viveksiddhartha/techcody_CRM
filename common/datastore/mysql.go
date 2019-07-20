package datastore

import (
	"SV_CRM/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"SV_CRM/common/utility"
)

func CreateUser(user *models.User) error {

	PasswordH := utility.SHA256OfString(user.PasswordHash)
	uuid := utility.GenerateUUID()

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO user(uuid, username, first_name, last_name, email, password_hash) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	fmt.Print("stmt details%v", stmt)

	defer stmt.Close()

	_, err = stmt.Exec(uuid, user.Username, user.FirstName, user.LastName, user.Email, PasswordH)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
