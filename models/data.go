
package models

import (
	"PMM/global"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	"context"
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

func SetUploadFile(params Data_set_upload_file)error{
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

	objectId, err := primitive.ObjectIDFromHex(params.ID)
    if err != nil {
		return err
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
		return err
	}	


	update := bson.D{{"$set", bson.D{
		{"tipefile1", params.Tipefile1},
    	{"tipefile2", params.Tipefile2},
    	{"tipefile3", params.Tipefile3},
    	{"tipefile4", params.Tipefile4},
    	{"urlfile1", fileimages["urlfile1"]},
    	{"urlfile2", fileimages["urlfile2"]},
    	{"urlfile3", fileimages["urlfile3"]},
    	{"urlfile4", fileimages["urlfile4"]}}}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Documents matched: %v\n", result.MatchedCount)
	fmt.Printf("Documents replaced: %v\n", result.ModifiedCount)

	return nil
}