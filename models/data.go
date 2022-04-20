
package models

import (
	"PMM/global"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	"context"
	// "errors"
	"strconv"
)

func Nasabah(CabangID,KelompokID string)([]bson.M,error) {
	var results []bson.M
	var filter bson.D

	fmt.Println(KelompokID)

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

	coll := client.Database("PMM").Collection("KELOMPOK")
	// filter := bson.D{{"cabangid", bson.D{{"$eq", OurBranchID}}}}
	filter := bson.D{{"cabangid", CabangID}}

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
	ID string `json:"id"`
	Tipefile1 string `json:"tipefile1"`
	Tipefile2 string `json:"tipefile2"`
	Tipefile3 string `json:"tipefile3"`
	Tipefile4 string `json:"tipefile4"`	
	Urlfile1 string `json:"urlfile1"`
	Urlfile2 string `json:"urlfile2"`
	Urlfile3 string `json:"urlfile3"`
	Urlfile4 string `json:"urlfile4"`
}

func SetUploadFile(params Data_set_upload_file)(Data_set_upload_file,error){
	var data Data_set_upload_file
	client,err := global.Conn_mongodb()
	if err != nil {
		return data,err
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()	
	coll := client.Database("PMM").Collection("NASABAH")

	objectId, err := primitive.ObjectIDFromHex(params.ID)
    if err != nil {
		return data,err
	}

	filter := bson.D{{"_id", objectId}}

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

	if params.Urlfile1 != "" {
		update := bson.D{{"$set", bson.D{
			{"tipefile1", params.Tipefile1},
			{"urlfile1", fileimages["urlfile1"]}}}}

		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return data,err
		}
		fmt.Printf("Documents matched: %v\n", result.MatchedCount)
		fmt.Printf("Documents replaced: %v\n", result.ModifiedCount)
	}
	if params.Urlfile2 != "" {
		update := bson.D{{"$set", bson.D{
			{"tipefile2", params.Tipefile2},
			{"urlfile2", fileimages["urlfile2"]}}}}

		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return data,err
		}
		fmt.Printf("Documents matched: %v\n", result.MatchedCount)
		fmt.Printf("Documents replaced: %v\n", result.ModifiedCount)
	}
	if params.Urlfile3 != "" {
		update := bson.D{{"$set", bson.D{
			{"tipefile3", params.Tipefile3},
			{"urlfile3", fileimages["urlfile3"]}}}}

		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return data,err
		}
		fmt.Printf("Documents matched: %v\n", result.MatchedCount)
		fmt.Printf("Documents replaced: %v\n", result.ModifiedCount)		
	}
	if params.Urlfile4 != ""{
		update := bson.D{{"$set", bson.D{
			{"tipefile4", params.Tipefile4},
			{"urlfile4", fileimages["urlfile4"]}}}}

		result, err := coll.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			return data,err
		}
		fmt.Printf("Documents matched: %v\n", result.MatchedCount)
		fmt.Printf("Documents replaced: %v\n", result.ModifiedCount)			
	}

	data = Data_set_upload_file{
		ID: params.ID,
		Tipefile1: params.Tipefile1,
		Tipefile2: params.Tipefile2,
		Tipefile3: params.Tipefile3,
		Tipefile4: params.Tipefile4,
		Urlfile1: fileimages["urlfile1"],
		Urlfile2: fileimages["urlfile2"],
		Urlfile3: fileimages["urlfile3"],
		Urlfile4: fileimages["urlfile4"],
	}


	return data,nil
}

func SetDownloadFile (CabangID,LoandID,Siklus string)error {
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

	coll := client.Database("PMM").Collection("NASABAH")

	Siklus_int, _ := strconv.Atoi(Siklus)
	
	filter = bson.D{{"cabangid", CabangID},{"loanid", LoandID},{"siklus", Siklus_int}}


	projection := bson.D{{"urlfile1", 1}, {"urlfile2", 1}, {"urlfile3", 1}, {"urlfile4", 1},{"_id", 0}}
	opts := options.Find().SetProjection(projection)
	
	cursor, err := coll.Find(context.TODO(),  filter, opts)
	if err != nil {
		return err
	}

	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}

	
	// response = Data_download{
	// 	Urlfile1: results[0]["urlfile1"].(string),
	// 	Urlfile2: results[0]["urlfile2"].(string),
	// 	Urlfile3: results[0]["urlfile3"].(string),
	// 	Urlfile4: results[0]["urlfile4"].(string),
	// }


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