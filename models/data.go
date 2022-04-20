
package models

import (
	"PMM/global"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	"context"
	// "errors"
	// "strconv"


	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	// "encoding/json"
	"time"
	"encoding/base64"
)

func Nasabah_bak(CabangID,KelompokID string)([]bson.M,error) {
	var results []bson.M
	var filter bson.D

	// fmt.Println(KelompokID)

	client,err := global.Conn_mongodb()
	if err != nil {
		return results,err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()	

	coll := client.Database("PMM").Collection("NASABAH")

	
	if CabangID != "" {
		filter = bson.D{{"cabangid", CabangID}}
	}
	if KelompokID != "null" {
		filter = bson.D{{"cabangid", CabangID},{"kelompokid", KelompokID}}
	}

	opts := options.Find()
	cursor, err := coll.Find(context.TODO(),  filter, opts)
	if err != nil {
		return results,err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return results,err
	}
	// for _, result := range results {
	// 	fmt.Println(result)
	// }

	return results,nil
}

func Nasabah(CabangID,KelompokID string)([]bson.M,error) {

	var results []bson.M
	var results2 []bson.M
	var filter bson.D
	var filter2 bson.D

	// fmt.Println(KelompokID)

	client,err := global.Conn_mongodb()
	if err != nil {
		return results,err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()	

	coll := client.Database("PMM").Collection("m_PNM_DW_BRNet_LastGet_DW")

	

	filter = bson.D{{"_id.CabangId", CabangID},{"KelompokId", KelompokID}}


	projection := bson.D{{"Nama", 1}, {"NasabahId", 1}, {"KelompokId", 1}, {"Produk", 1},{"_id", 1}}
	opts := options.Find().SetProjection(projection)

	// opts := options.Find()
	cursor, err := coll.Find(context.TODO(),  filter, opts)
	if err != nil {
		return results,err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return results,err
	}


	for _, result := range results {

		
		Pcabangid := result["_id"].(bson.M)["CabangId"].(string)
		Ploandid := result["_id"].(bson.M)["LoanId"].(string)
		Psiklus := fmt.Sprint(result["_id"].(bson.M)["Siklus"].(int32))		

		coll2 := client.Database("PMM").Collection("transaksi_dokumenupload")

		filter2 = bson.D{{"cabangid", Pcabangid},{"loanid", Ploandid},{"siklus", Psiklus}}


		opts2 := options.Find()
		cursor2, err := coll2.Find(context.TODO(),  filter2, opts2)
		if err != nil {
			return results,err
		}
	
		if err = cursor2.All(context.TODO(), &results2); err != nil {
			return results,err
		}		


		// fmt.Println(results2)
		result["_id"] = base64.StdEncoding.EncodeToString([]byte(Pcabangid+"-"+Ploandid+"-"+Psiklus))
		result["cabangid"] = Pcabangid
		result["loanid"] = Ploandid
		result["siklus"] = Psiklus
		result["tipefile1"] = ""
		result["tipefile2"] = ""
		result["tipefile3"] = ""
		result["tipefile4"] = ""
		result["urlfile1"] = ""
		result["urlfile2"] = ""
		result["urlfile3"] = ""
		result["urlfile4"] = ""

		for _, result2 := range results2 {
			if (Pcabangid == result2["cabangid"] && Ploandid == result2["loanid"] && Psiklus == result2["siklus"]){
				result["tipefile1"] = result2["tipefile1"]
				result["tipefile2"] = result2["tipefile2"]
				result["tipefile3"] = result2["tipefile3"]
				result["tipefile4"] = result2["tipefile4"]
				result["urlfile1"] = result2["urlfile1"]
				result["urlfile2"] = result2["urlfile2"]
				result["urlfile3"] = result2["urlfile3"]
				result["urlfile4"] = result2["urlfile4"]
			}
		}	
	}

	return results,nil
}

func Kelompok(CabangID string)([]bson.M,error) {
	var results []bson.M

	client,err := global.Conn_mongodb()
	if err != nil {
		return results,err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()	

	coll := client.Database("PMM").Collection("master_kelompok")
	// filter := bson.D{{"cabangid", bson.D{{"$eq", OurBranchID}}}}
	filter := bson.D{{"CabangId", CabangID}}

	opts := options.Find()
	cursor, err := coll.Find(context.TODO(),  filter, opts)
	if err != nil {
		return results,err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return results,err
	}
	// for _, result := range results {
	// 	fmt.Println(result)
	// }

	return results,nil
}

type Data_set_upload_file struct {
	// ID string `json:"id"`
	Loandid string `json:"loanid"`
	Cabangid string `json:"cabangid"`
	Siklus string `json:"siklus"`
	Tipefile1 string `json:"tipefile1"`
	Tipefile2 string `json:"tipefile2"`
	Tipefile3 string `json:"tipefile3"`
	Tipefile4 string `json:"tipefile4"`	
	Urlfile1 string `json:"urlfile1"`
	Urlfile2 string `json:"urlfile2"`
	Urlfile3 string `json:"urlfile3"`
	Urlfile4 string `json:"urlfile4"`
	Created_by string `json:"created_by"`
}

func SetUploadFile(params Data_set_upload_file)(map[string]string,error){
	// var data Data_set_upload_file
	var check_data []bson.M
	var data = make(map[string]string)

	client,err := global.Conn_mongodb()
	if err != nil {
		return data,err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()	
	coll := client.Database("PMM").Collection("transaksi_dokumenupload_new")

	//CHECK DATA
	filter := bson.D{{"cabangid", params.Cabangid},{"loanid", params.Loandid},{"siklus", params.Siklus}}


	opts := options.Find()
	cursor, err := coll.Find(context.TODO(),  filter, opts)
	if err != nil {
		return data,err
	}

	if err = cursor.All(context.TODO(), &check_data); err != nil {
		return data,err
	}




	//SIMPAN FILE
	var fileimages = make(map[string]string)
	fileimages["urlfile1"] = params.Urlfile1
	fileimages["urlfile2"] = params.Urlfile2
	fileimages["urlfile3"] = params.Urlfile3
	fileimages["urlfile4"] = params.Urlfile4	
	fileimages,err = global.MapB64SaveFile(fileimages,"/")
	if err != nil { 			
		return data,err
	}	

	// fmt.Println(check_data)
	if len(check_data) > 0 {
		//UPDATE DATA
		filter := bson.D{{"cabangid", params.Cabangid},{"loanid", params.Loandid},{"siklus", params.Siklus}}
		update := bson.D{{"$set", bson.D{
						{"tipefile1", params.Tipefile1},
						{"tipefile2", params.Tipefile2},
						{"tipefile3", params.Tipefile3},
						{"tipefile4", params.Tipefile4},
						{"urlfile1", fileimages["urlfile1"]},
						{"urlfile2", fileimages["urlfile2"]},
						{"urlfile3", fileimages["urlfile3"]},
						{"urlfile4", fileimages["urlfile3"]},
						{"modified_by", params.Created_by},
						{"modified_time", time.Now().Format("2006-01-02 15:04:05")}}}}

		result, err := coll.UpdateOne(context.TODO(), filter, update)	
		if err != nil { 			
			return data,err
		}	
		fmt.Printf("Documents matched: %v\n", result.MatchedCount)
		fmt.Printf("Documents updated: %v\n", result.ModifiedCount)


	} else {
		//INSERT DATA
		doc := bson.D{{"loanid", params.Loandid},
					{"cabangid", params.Cabangid}, 
					{"siklus", params.Siklus}, 
					{"tipefile1", params.Tipefile1}, 
					{"tipefile2", params.Tipefile2}, 
					{"tipefile3", params.Tipefile3}, 
					{"tipefile4", params.Tipefile4}, 
					{"urlfile1", fileimages["urlfile1"]}, 
					{"urlfile2", fileimages["urlfile2"]}, 
					{"urlfile3", fileimages["urlfile3"]}, 
					{"urlfile4", fileimages["urlfile4"]}, 
					{"created_by", params.Created_by}, 
					{"created_time", time.Now().Format("2006-01-02 15:04:05")}, 
					{"modified_by", ""}, 
					{"modified_time", ""}}


		result, err := coll.InsertOne(context.TODO(), doc)
		if err != nil {
			return data,err
		}
		
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	}

	data["loanid"] = params.Loandid
	data["cabangid"] = params.Cabangid
	data["siklus"] = params.Siklus
	data["tipefile1"] = params.Tipefile1
	data["tipefile2"] = params.Tipefile2
	data["tipefile3"] = params.Tipefile3
	data["tipefile4"] = params.Tipefile4
	data["urlfile1"] = fileimages["urlfile1"]
	data["urlfile2"] = fileimages["urlfile2"]
	data["urlfile3"] = fileimages["urlfile3"]
	data["urlfile4"] = fileimages["urlfile4"]

	return data,nil
}

func SetDownloadFile(CabangID,LoandID,Siklus string)error {
	var results []bson.M
	var filter bson.D

	client,err := global.Conn_mongodb()
	if err != nil {
		return err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()	

	coll := client.Database("PMM").Collection("transaksi_dokumenupload_new")

	// Siklus_int, _ := strconv.Atoi(Siklus)
	
	filter = bson.D{{"cabangid", CabangID},{"loanid", LoandID},{"siklus", Siklus}}


	projection := bson.D{{"urlfile1", 1}, {"urlfile2", 1}, {"urlfile3", 1}, {"urlfile4", 1},{"_id", 0}}
	opts := options.Find().SetProjection(projection)
	
	cursor, err := coll.Find(context.TODO(),  filter, opts)
	if err != nil {
		return err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}

	fmt.Println(results)


	err = global.GetSavetoPathS3(results[0]["urlfile1"].(string),"./temp/"+CabangID+LoandID+Siklus) 
	if err != nil {return err}	
	err = global.GetSavetoPathS3(results[0]["urlfile2"].(string),"./temp/"+CabangID+LoandID+Siklus) 
	if err != nil {return err}	
	err = global.GetSavetoPathS3(results[0]["urlfile3"].(string),"./temp/"+CabangID+LoandID+Siklus) 
	if err != nil {return err}	
	err = global.GetSavetoPathS3(results[0]["urlfile4"].(string),"./temp/"+CabangID+LoandID+Siklus) 
	if err != nil {return err}	


	err = global.ZipWriter("./temp/"+CabangID+LoandID+Siklus,"./temp/data"+CabangID+LoandID+Siklus)
	if err!=nil {
		return err
	}

	return nil
}




