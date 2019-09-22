package datastore

import (
	"fmt"
	"log"
	"svcrm/models"

	_ "github.com/go-sql-driver/mysql"

	"svcrm/common/utility"
)

func EntityCreate(Entity *models.CoEntity) error {

	uuid := utility.GenerateUUID()
	SecretKey := utility.SHA256OfString(uuid)
	Password := utility.SHA256OfString(Entity.Password)
	Country := "INDIA"

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	fmt.Println(" values in the string %v & %v & %v & %v & %v & %v & %v & %v", uuid, Entity.CoEntityId, Entity.CompanyNm, Entity.AliasNm, Entity.State, Country, Entity.Email, SecretKey)

	stmt, err := tx.Prepare("INSERT INTO coentity(uuid, CoEntityId, CompanyNm, AliasNm, State, Country, Email, SecretKey, Password) VALUES (?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uuid, Entity.CoEntityId, Entity.CompanyNm, Entity.AliasNm, Entity.State, Country, Entity.Email, SecretKey, Password)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
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

func CreateContract(con *models.Contract) error {

	uuid := utility.GenerateUUID()
	SecretKey := utility.SHA256OfString(uuid)
	JsonBlock := utility.SHA256OfString(SecretKey)
	//EffectiveDate, err := strconv.Atoi(con.EffectiveDate)

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	fmt.Println(" values in the string %v & %v & %v & %v & %v & %v & %v \n", uuid, con.CoEntityID, con.Version, con.EffectiveDate, con.ContractType, con.JsonObject, JsonBlock, SecretKey)

	stmt, err := tx.Prepare("INSERT INTO contracts(ContractID, CoEntityId, Version, EffectiveDate, ContractType, JsonObject, JsonBlock, SecretKey) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uuid, con.CoEntityID, con.Version, con.EffectiveDate, con.ContractType, con.JsonObject, JsonBlock, SecretKey)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err

}

//TODO: Need to append lenth logic
func CreateAllocation(con *models.Contract) error {

	uuid := utility.GenerateUUID()

	Contract, err := GetContractsByCoEntityID(con.CoEntityID)

	contractID := Contract.ContractID

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO allocation (uuid, ContractID, CoEntityID, ProfileID, Allocation, ContractType, Relation, JsonObject) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	for i := 0; i < len(con.Allocation); {
		_, err = stmt.Exec(uuid, contractID, con.CoEntityID, con.Allocation[i].ProfileName, con.Allocation[i].Percentage, con.ContractType, con.Allocation[i].Relation, con.JsonObject)
		if err != nil {
			return err
		}
		i++
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	return err

}

//===Get queries=======================================================

func GetProfileDetailsByCoEntityProfilename(CoEntityID string, Profilename string) (*models.Profile, error) {

	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, profilename, first_name, last_name, email,EmailVerified, password_hash,ContactNo,PhoneVerified,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM Profile WHERE Status in ('0','1') and CoEntityId=? and Profilename = ? ")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(CoEntityID, Profilename)
	u := models.Profile{}
	err = row.Scan(&u.UUID, &u.CoEntityID, &u.Profilename, &u.FirstName, &u.LastName, &u.Email, &u.EmailVerified, &u.PasswordHash, &u.ContactNo, &u.PhoneVerified, &u.Status, &u.TimestampCreated, &u.TimestampModified)
	return &u, err

}

func GetProfileDetailsByProfil(Profilename string) (*models.Profile, error) {

	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, profilename, first_name, last_name, email,EmailVerified, password_hash,ContactNo,PhoneVerified,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM Profile WHERE Status in ('0','1') and Profilename = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(Profilename)
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

	stmt, err := m.Prepare("SELECT uuid, CoEntityID, CompanyNm, AliasNm, State, Country,Email, Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM coentity WHERE Status in ('0','1') and CoEntityID = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(CoEntityId)
	u := models.CoEntity{}
	err = row.Scan(&u.UUID, &u.CoEntityId, &u.CompanyNm, &u.AliasNm, &u.State, &u.Country, &u.Email, &u.Status, &u.TimestampCreated, &u.TimestampModified)
	return &u, err
}
func GetEntityDetailsByCoEntityIdForPassword(CoEntityId string) (*models.CoEntity, error) {
	m := DBConn()

	stmt, err := m.Prepare("SELECT uuid, CoEntityID, CompanyNm, AliasNm, State, Country,Email, SecretKey, password, Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM coentity WHERE Status in ('0','1') and CoEntityID = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(CoEntityId)
	u := models.CoEntity{}
	err = row.Scan(&u.UUID, &u.CoEntityId, &u.CompanyNm, &u.AliasNm, &u.State, &u.Country, &u.Email, &u.SecretKey, &u.Password, &u.Status, &u.TimestampCreated, &u.TimestampModified)
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

func GetProfileDetailsWithoutStatusByContactNo(ContactNo string) (*models.Profile, error) {

	m := DBConn()
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, profilename, first_name, last_name, email,EmailVerified, password_hash,ContactNo,PhoneVerified,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM Profile WHERE ContactNo = ?")
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
func GetProfileDetailsWithoutStatusByemail(email string) (*models.Profile, error) {

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

func GetContractsByCoEntityID(CoEntityID string) (*models.Contract, error) {

	m := DBConn()
	stmt, err := m.Prepare("SELECT ContractID, CoEntityID, Version, EffectiveDate, ContractType, Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM contracts WHERE Status in ('0','1') and CoEntityID = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(CoEntityID)
	u := models.Contract{}
	err = row.Scan(&u.ContractID, &u.CoEntityID, &u.Version, &u.EffectiveDate, &u.ContractType, &u.Status, &u.TimestampCreated, &u.TimestampModified)
	return &u, err

}

//===Update queries=======================================================

func UpdateProfileByProfileID(Profile *models.Profile) error {

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("UPDATE Profile set first_name =?, last_name =?, email =?,ContactNo =? where Profilename = ?")

	if err != nil {
		return err
	}
	fmt.Print("stmt details%v", stmt)

	defer stmt.Close()

	_, err = stmt.Exec(Profile.FirstName, Profile.LastName, Profile.Email, Profile.ContactNo, Profile.Profilename)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateEntityByEntityID(Entity *models.CoEntity) error {

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("UPDATE CoEntity SET CompanyNm =?, AliasNm =?, State =?, Email =?  where CoEntityId=?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(Entity.CompanyNm, Entity.AliasNm, Entity.State, Entity.Email, Entity.CoEntityId)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateContract(con *models.Contract) error {

	uuid := utility.GenerateUUID()
	SecretKey := utility.SHA256OfString(uuid)
	JsonBlock := utility.SHA256OfString(SecretKey)
	//EffectiveDate, err := strconv.Atoi(con.EffectiveDate)

	m := DBConn()
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	fmt.Println(" values in the string %v & %v & %v & %v & %v & %v & %v \n", uuid, con.CoEntityID, con.Version, con.EffectiveDate, con.ContractType, con.JsonObject, JsonBlock, SecretKey)

	stmt, err := tx.Prepare("UPDATE contracts SET (ContractID, CoEntityId, Version, EffectiveDate, ContractType, JsonObject, JsonBlock, SecretKey) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(uuid, con.CoEntityID, con.Version, con.EffectiveDate, con.ContractType, con.JsonObject, JsonBlock, SecretKey)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return err

}

//=========== GET multiple record List

func GetAllProfileDetailsByCoEntity(CoEntityId string) ([]models.Profile, error) {

	m := DBConn()
	Profile := make([]models.Profile, 0)
	stmt, err := m.Prepare("SELECT uuid, CoEntityID, profilename, first_name, last_name, email,EmailVerified, ContactNo,PhoneVerified,Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM Profile WHERE Status in ('0','1') and CoEntityId= ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(CoEntityId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := models.Profile{}
		err := rows.Scan(&u.UUID, &u.CoEntityID, &u.Profilename, &u.FirstName, &u.LastName, &u.Email, &u.EmailVerified, &u.ContactNo, &u.PhoneVerified, &u.Status, &u.TimestampCreated, &u.TimestampModified)
		if err != nil {
			return nil, err
		}
		Profile = append(Profile, u)
	}
	return Profile, nil
}

func GetAllContractByCoEntity(CoEntityId string) ([]models.Contract, error) {

	m := DBConn()
	Contract := make([]models.Contract, 0)
	stmt, err := m.Prepare("SELECT ContractID, CoEntityID, Version, EffectiveDate, ContractType, Status, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM contracts WHERE Status in ('0','1') and CoEntityID = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(CoEntityId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := models.Contract{}
		err := rows.Scan(&u.ContractID, &u.CoEntityID, &u.Version, &u.EffectiveDate, &u.ContractType, &u.Status, &u.TimestampCreated, &u.TimestampModified)
		if err != nil {
			return nil, err
		}
		Contract = append(Contract, u)
	}
	return Contract, nil
}

func GetAllocationByCoEntityContractID(CoEntityID string) ([]models.AllocationList, error) {

	m := DBConn()
	AllocationList := make([]models.AllocationList, 0)
	contractId, err := GetContractsByCoEntityID(CoEntityID)
	if err != nil {
		return nil, err
	}

	stmt, err := m.Prepare("SELECT allocation.uuid, contracts.CoEntityID, contracts.ContractID, allocation.ProfileID, contracts.ContractType, contracts.EffectiveDate, allocation.Allocation,  allocation.Relation, allocation.status, contracts.Version, allocation.created_ts, allocation.updated_ts FROM contracts , allocation WHERE contracts.ContractID = allocation.ContractID AND allocation.Status in ('0','1') and contracts.CoEntityID = ? and contracts.ContractID=?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(CoEntityID, contractId.ContractID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		u := models.AllocationList{}
		err := rows.Scan(&u.UUID, &u.CoEntityID, &u.ContractID, &u.ProfileName, &u.ContractType, &u.EffectiveDate, &u.Percentage, &u.Relation, &u.Status, &u.Version, &u.TimestampCreated, &u.TimestampModified)
		if err != nil {
			return nil, err
		}
		AllocationList = append(AllocationList, u)
	}

	return AllocationList, nil
}
