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

	stmt, err := tx.Prepare("INSERT INTO Profile(uuid, CoEntityID, Profilename, first_name, last_name, email,ContactNo, password_hash) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	fmt.Print("stmt details%v", stmt)

	defer stmt.Close()

	_, err = stmt.Exec(uuid, Profile.CoEntityID, Profile.Profilename, Profile.FirstName, Profile.LastName, Profile.Email, Profile.ContactNo, PasswordH)
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

func GetProfileDetailsByProfilename(Profilename string, CoEntityID string) (*models.Profile, error) {

	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, profilename, first_name, last_name, email,EmailVerified, password_hash,ContactNo,PhoneVerified,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM Profile WHERE Status in ('0','1') and Profilename = ? and CoEntityID=?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(Profilename, CoEntityID)
	u := models.Profile{}
	err = row.Scan(&u.UUID, &u.CoEntityID, &u.Profilename, &u.FirstName, &u.LastName, &u.Email, &u.EmailVerified, &u.PasswordHash, &u.ContactNo, &u.PhoneVerified, &u.Status, &u.TimestampCreated, &u.TimestampModified)
	return &u, err

}
func GetProfileDetailsByCoEntity(CoEntityId string) (*models.Profile, error) {

	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, profilename, first_name, last_name, email,EmailVerified, password_hash,ContactNo,PhoneVerified,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM Profile WHERE Status in ('0','1') and CoEntityId= ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(CoEntityId)
	u := models.Profile{}
	err = row.Scan(&u.UUID, &u.CoEntityID, &u.Profilename, &u.FirstName, &u.LastName, &u.Email, &u.EmailVerified, &u.PasswordHash, &u.ContactNo, &u.PhoneVerified, &u.Status, &u.TimestampCreated, &u.TimestampModified)
	return &u, err

}
func GetProfileDetailsByemail(email string) (*models.Profile, error) {

	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, profilename, first_name, last_name, email,EmailVerified, password_hash,ContactNo,PhoneVerified,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM Profile WHERE Status in ('0','1') and email = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)
	u := models.Profile{}
	err = row.Scan(&u.UUID, &u.CoEntityID, &u.Profilename, &u.FirstName, &u.LastName, &u.Email, &u.EmailVerified, &u.PasswordHash, &u.ContactNo, &u.PhoneVerified, &u.Status, &u.TimestampCreated, &u.TimestampModified)
	return &u, err

}

func GetProfileDetailsByContactNo(ContactNo string) (*models.Profile, error) {

	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, profilename, first_name, last_name, email,EmailVerified, password_hash,ContactNo,PhoneVerified,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM Profile WHERE Status in ('0','1') and ContactNo = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(ContactNo)
	u := models.Profile{}
	err = row.Scan(&u.UUID, &u.CoEntityID, &u.Profilename, &u.FirstName, &u.LastName, &u.Email, &u.EmailVerified, &u.PasswordHash, &u.ContactNo, &u.PhoneVerified, &u.Status, &u.TimestampCreated, &u.TimestampModified)
	return &u, err

}

func GetEntityDetailsByCoEntityId(CoEntityId string) (*models.CoEntity, error) {
	m := DBConn()

	stmt, err := m.Prepare("SELECT uuid, CoEntityID, CompanyNm, AliasNm, State, Country,Email, SecretKey,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM coentity WHERE Status in ('0','1') and CoEntityID = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(CoEntityId)
	u := models.CoEntity{}
	err = row.Scan(&u.UUID, &u.CoEntityId, &u.CompanyNm, &u.AliasNm, &u.SecretKey, &u.Country, &u.Email, &u.SecretKey, &u.TimestampCreated, &u.TimestampModified)
	return &u, err
}

func GetEntityDetailsByCompanyNm(CompanyNm string) (*models.CoEntity, error) {
	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, CompanyNm, AliasNm, State, Country,Email, SecretKey,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM coentity WHERE Status in ('0','1') and CompanyNm = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(CompanyNm)
	u := models.CoEntity{}
	err = row.Scan(&u.UUID, &u.CoEntityId, &u.CompanyNm, &u.AliasNm, &u.SecretKey, &u.Country, &u.Email, &u.SecretKey, &u.TimestampCreated, &u.TimestampModified)
	return &u, err
}

func GetEntityDetailsByEmail(Email string) (*models.CoEntity, error) {
	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, CompanyNm, AliasNm, State, Country,Email, SecretKey,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM coentity WHERE Status in ('0','1') and Email = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(Email)
	u := models.CoEntity{}
	err = row.Scan(&u.UUID, &u.CoEntityId, &u.CompanyNm, &u.AliasNm, &u.SecretKey, &u.Country, &u.Email, &u.SecretKey, &u.TimestampCreated, &u.TimestampModified)
	return &u, err
}

/*
func GetEntityDetailsByGenericParam(u *models.CoEntity) {
	m := DBConn()

	row := m.QueryRow("SELECT uuid, CoEntityID, CompanyNm, AliasNm, State, Country,Email, SecretKey,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM coentity WHERE Status in ('0','1') and CoEntityID = $1", u.CoEntityId)

	err = row.Scan(&u.UUID, &u.CoEntityId, &u.CompanyNm, &u.AliasNm, &u.SecretKey, &u.Country, &u.Email, &u.SecretKey, &u.TimestampCreated, &u.TimestampModified)
	return nil, err

}
*/
