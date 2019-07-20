package datastore

import (
	"SV_CRM/models"
	"database/sql"
)

type Datastore interface {
	CreateProfile(Profile *models.Profile) error
	EntityCreate(Partner *models.CoEntity)
	Close()
}

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, "sv_crm:sv_crm@tcp(127.0.0.1:3306)/gopherfacedb")
	if err != nil {
		panic(err.Error())
	}
	return db

}
