
package models

import (
	"PMM/global"
	// "fmt"
)


type Master_cabang struct {
	OurBranchID string `json:"ourbranchid" db:"ourbranchid"`
	Branchname string `json:"branchname" db:"branchname"`
}

func Cabang()([]Master_cabang,error) {
	data := make([]Master_cabang, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()	
	var query string


	query = " SELECT ourbranchid,branchname FROM OPENQUERY([10.61.3.15],'select wilayah, Regionname, areaname, ourbranchid,branchname from MKR_DATA.dbo.Cabang_Mekaar')  "

	err = db.RawSQL(query).Do(&data)

	if err != nil {		
		// global.Logging("ERROR","models GetIdProspekTerpakai ---> "+query+" ---> " + err.Error())	
		return data,err
	}		

	return data,nil
}

type Master_tipefile struct {
	Id string `json:"id" db:"id"`
	Tipefile string `json:"tipefile" db:"tipefile"`
}

func Tipefile()([]Master_tipefile,error) {
	data := make([]Master_tipefile, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()	
	var query string


	query = " SELECT id,tipefile from master_tipefile"

	err = db.RawSQL(query).Do(&data)

	if err != nil {		
		// global.Logging("ERROR","models GetIdProspekTerpakai ---> "+query+" ---> " + err.Error())	
		return data,err
	}		

	return data,nil
}

type Master_role struct {
	Id_role string `json:"id_role" db:"id_role"`
	Nama_role string `json:"nama_role" db:"nama_role"`
	Isactive string `json:"isactive" db:"isactive"`
}

func Role()([]Master_role,error) {
	data := make([]Master_role, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()	
	var query string


	query = " SELECT id_role,nama_role,isactive from master_role"

	err = db.RawSQL(query).Do(&data)

	if err != nil {		
		// global.Logging("ERROR","models GetIdProspekTerpakai ---> "+query+" ---> " + err.Error())	
		return data,err
	}		

	return data,nil
}

type Master_role_detail struct {
	Id string `json:"id" db:"id"`
	Id_role string `json:"id_role" db:"id_role"`
	Menu string `json:"menu" db:"menu"`
	Submenu string `json:"submenu" db:"submenu"`
	Create string `json:"create" db:"create"`
	Read string `json:"read" db:"read"`
	Update string `json:"update" db:"update"`
	Delete string `json:"delete" db:"delete"`
}

func RoleDetail()([]Master_role_detail,error) {
	data := make([]Master_role_detail, 0)
	db,err := global.Conn()
	if err != nil {
		return data,err
	}		
	defer db.Close()	
	var query string


	query = " SELECT [id],[id_role],[menu],[submenu],[create],[read],[update],[delete] from master_roledetil"

	err = db.RawSQL(query).Do(&data)

	if err != nil {		
		// global.Logging("ERROR","models GetIdProspekTerpakai ---> "+query+" ---> " + err.Error())	
		return data,err
	}		

	return data,nil
}