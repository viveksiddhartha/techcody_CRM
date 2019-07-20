package datastore

import (
	"SV_CRM/models"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"SV_CRM/common/utility"
)

func CreateProfile(Profile *models.Profile) error {

	PasswordH := utility.SHA256OfString(Profile.PasswordHash)
	uuid := utility.GenerateUUID()

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO Profile(uuid, CoEntity, Profilename, first_name, last_name, email,ContactNo, password_hash) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	fmt.Print("stmt details%v", stmt)

	defer stmt.Close()

	_, err = stmt.Exec(uuid, Profile.CoEntity, Profile.Profilename, Profile.FirstName, Profile.LastName, Profile.Email, Profile.ContactNo, PasswordH)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func EntityCreate(Entity *models.CoEntity) error {

	uuid := utility.GenerateUUID()
	SecretKey := utility.SHA256OfString(uuid)
	Country := "INDIA"

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	fmt.Println(" values in the string %v & %v & %v & %v & %v & %v & %v & %v", uuid, Entity.CoEntityId, Entity.CompanyNm, Entity.AliasNm, Entity.State, Country, Entity.Email, SecretKey)

	stmt, err := tx.Prepare("INSERT INTO CoEntity(uuid, CoEntityId, CompanyNm, AliasNm, State, Country, Email, SecretKey) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uuid, Entity.CoEntityId, Entity.CompanyNm, Entity.AliasNm, Entity.State, Country, Entity.Email, SecretKey)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
