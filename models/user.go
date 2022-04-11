package models

import (
	"PMM/global"
	// "fmt"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "context"
	"crypto/sha1"
	"crypto/md5"
	"encoding/hex"
	b64 "encoding/base64"
	// "fmt"
	"errors"
)

type T_master_user struct {
	Username string `json:"username" db:"username"`
	Nip	string `json:"nip" db:"nip"`
	Nama string `json:"nama" db:"nama"`
	Id_role string `json:"id_role" db:"id_role"`
	Cabang string `json:"cabang" db:"cabang"`
	Unit_kerja string `json:"unit_kerja" db:"unit_kerja"`
	Lokasi string `json:"lokasi" db:"lokasi"`	
	Token string `json:"string_token" `
}

func Login_mekaar_integrasi(user,password string)([]T_master_user,error) {
	data := make([]T_master_user, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()	
	var query string

	md5 := md5.Sum([]byte(password))
	sha1 := sha1.Sum([]byte(hex.EncodeToString(md5[:])))

	password = hex.EncodeToString(sha1[:])

	// fmt.Println(hex.EncodeToString(md5[:]))
	// fmt.Println(hex.EncodeToString(sha1[:]))

	query = "SELECT username,nip,nama,id_role,cabang,ISNULL(unit_kerja,'') as unit_kerja,lokasi FROM master_user WHERE CONCAT(username,password)=CONCAT('"+user+"','"+password+"') and lokasi='0' "	

	err = db.RawSQL(query).Do(&data)

	if err != nil {		
		// global.Logging("ERROR","models GetIdProspekTerpakai ---> "+query+" ---> " + err.Error())	
		return data,err
	}		

	if len(data) < 1 {
		return data,errors.New("user / password salah")
	}

	data[0].Token = global.GenerateToken(password)

	return data,nil
}

func Login_sso(user,password string)([]T_master_user,error) {
	data := make([]T_master_user, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()	
	var query string

	password = b64.StdEncoding.EncodeToString([]byte(password))


	query = " SELECT username,nip,nama,id_role,ISNULL(cabang,'') as cabang,ISNULL(unit_kerja,'') as unit_kerja,lokasi FROM master_user WHERE CONCAT(username,password)=CONCAT('"+user+"','"+password+"') and lokasi='1'  "

	err = db.RawSQL(query).Do(&data)

	if err != nil {		
		// global.Logging("ERROR","models GetIdProspekTerpakai ---> "+query+" ---> " + err.Error())	
		return data,err
	}		

	if len(data) < 1 {
		return data,errors.New("user / password salah")
	}

	data[0].Token = global.GenerateToken(password)

	return data,nil
}