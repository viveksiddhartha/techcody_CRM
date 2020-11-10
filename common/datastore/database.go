package datastore

import (
	"database/sql"
	"log"
	"svcrm/models"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore interface {
	CreateProfile(Profile *models.Profile) error
	EntityCreate(Entity *models.CoEntity)
	CreateContract(con *models.Allocation) (*models.Allocation, error)
	CreateAllocation(con *models.Allocation) (*models.Allocation, error)
	Close()

	//========GET ENTITY Details
	GetEntityDetailsByCoEntityId(CoEntityId string) (*models.CoEntity, error)
	GetEntityDetailsByCompanyNm(CompanyNm string) (*models.CoEntity, error)
	GetEntityDetailsByEmail(Email string) (*models.CoEntity, error)
	GetEntityDetailsByCoEntityIdForPassword(CoEntityId string) (*models.CoEntity, error)

	//========GET Profile Procedure
	GetProfileDetailsByCoEntityProfilename(CoEntityID string, Profilename string) (*models.Profile, error)
	GetProfileDetailsByProfilename(Profilename string) (*models.Profile, error)
	GetProfileDetailsByCoEntity(CoEntity string) (*models.Profile, error)
	GetProfileDetailsByemail(email string) (*models.Profile, error)
	GetProfileDetailsByContactNo(ContactNo string) (*models.Profile, error)
	GetProfileDetailsWithoutStatusByemail(email string) (*models.Profile, error)
	GetProfileDetailsWithoutStatusByContactNo(ContactNo string) (*models.Profile, error)
	GetProfileDetailsByProfil(Profilename string) (*models.Profile, error)

	//======GET Contract ===================
	GetContractsByCoEntityID(CoEntityID string) (*models.Allocation, error)

	//==========UpdateProfile
	UpdateProfileByProfileID(Profile *models.Profile) error
	UpdateEntityByEntityID(Entity *models.CoEntity) error
	UpdateContract(con *models.Contract) error

	//==========GET All
	GetAllProfileDetailsByCoEntity(CoEntityId string) ([]models.Profile, error)
	GetAllocationByCoEntityContractID(CoEntityId string) ([]models.AllocationList, error)
	GetAllContractByCoEntity(CoEntityId string) ([]models.Contract, error)
}

type RDatastore interface {
	CreateEntityRedis(entity *models.CoEntity) error
	GetUserRedis(username string) (*models.CoEntity, error)
}

//var db sql.DB

func DBConn() (db *sql.DB) {
	dbDriver := "mysql"
	db, err := sql.Open(dbDriver, "svcrm:Pass#word1@tcp(localhost:3306)/sv_crm")
	db.SetConnMaxLifetime(20)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(5)
	db.Stats()
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		panic(err.Error())
	}

	return db

}

/*
func RDBConn() (dbs *sql.DB) {
	dbDriver := "redis"
	dbs, err := sql.Open(dbDriver, "sv_crm:sv_crm@tcp(127.0.0.1:3306)/SV_CRM")
	if err != nil {
		panic(err.Error())
	}
	return dbs

}
*/
