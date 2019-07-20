package datastore

import (
	"SV_CRM/models"
	"database/sql"
	"fmt"
)

type Datastore interface {
	CreateUser(user *models.User) error
	Close()
}

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, "sv_crm:sv_crm@tcp(127.0.0.1:3306)/gopherfacedb")
	fmt.Println("DB connection successful")
	if err != nil {
		panic(err.Error())
	}
	return db

}
