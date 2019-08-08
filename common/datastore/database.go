package datastore

import (
	"SV_CRM/models"
	"database/sql"
)

type Datastore interface {
	CreateProfile(Profile *models.Profile) error
	EntityCreate(Entity *models.CoEntity)
	Close()
	GetProfileDetailsByProfilename(Profilename string) (*models.Profile, error)
	GetProfileDetailsByCoEntity(CoEntity string) (*models.Profile, error)
	GetProfileDetailsByemail(email string) (*models.Profile, error)
	GetProfileDetailsByContactNo(ContactNo string) (*models.Profile, error)
	GetEntityDetailsByCoEntityId(CoEntityId string) (*models.CoEntity, error)
	GetEntityDetailsByCompanyNm(CompanyNm string) (*models.CoEntity, error)
	GetEntityDetailsByEmail(Email string) (*models.CoEntity, error)
	UpdateProfileByProfileID(Profile *models.Profile) error
	GetProfileDetailsByProfilCoEntity(Profilename string, CoEntityID string) (*models.Profile, error)
	GetProfileDetailsWithoutStatusByemail(email string) (*models.Profile, error)
	GetProfileDetailsWithoutStatusByContactNo(ContactNo string) (*models.Profile, error)
	UpdateEntityByEntityID(Entity *models.CoEntity) error
}

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, "sv_crm:sv_crm@tcp(127.0.0.1:3306)/SV_CRM")
	if err != nil {
		panic(err.Error())
	}
	return db

}
