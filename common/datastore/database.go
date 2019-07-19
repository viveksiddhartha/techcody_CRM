package datastore

import (
	"SV_CRM/models"
	"database/sql"
	"fmt"
)

type Datastore interface {
	CreateUser(user *models.User) error
}

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, "root:toortoor@tcp(127.0.0.1:3306)/gopherfacedb")
	fmt.Println("DB connection successful")
	if err != nil {
		panic(err.Error())
	}
	return db

}
