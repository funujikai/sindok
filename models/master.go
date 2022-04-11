
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