package global

import (
	beego "github.com/beego/beego/v2/server/web"
	"fmt"
	"strconv"
	"time"	
	"crypto/md5"
	"github.com/dgrijalva/jwt-go"	
	"io"
	"log"		

	"github.com/minio/minio-go/v7"
	"golang.org/x/net/context"
	"os"
	"encoding/base64"
	"strings"    
    "io/ioutil"
	"github.com/beevik/guid"
	// "image"
	"bytes"
    _ "image/gif"
    _ "image/jpeg"
    _ "image/png"   	
	"errors"
)


func GenerateToken(d string) string {
	var uid int = 0

	KEY,_ := beego.AppConfig.String("KEY")

	currentTimestamp := time.Now().UTC().Unix()
	//var ttl int64 = 3600 //satu jam
	var ttl int64 = 43200 //dua belas jam
	h := md5.New()
	_,err := io.WriteString(h, strconv.Itoa(uid))

	if err != nil {
    	log.Fatal(err)
	}

	_,err = io.WriteString(h, strconv.FormatInt(int64(currentTimestamp), 10))

	if err != nil {
    	log.Fatal(err)
	}	

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": uid,
		"iat": currentTimestamp,
		"exp": currentTimestamp + ttl,
		"nbf": currentTimestamp,
		"iss": d,
		"jti": h.Sum(nil),
	})

	tokenString, err := token.SignedString([]byte(KEY))

	if err != nil {
    	log.Fatal(err)
	}

	return (tokenString)
}



func PutS3(file_src, file_s3 string)error{
    ctx := context.Background()
    minioClient,err := ConnS3Storage()
    if err != nil {
        return err
    }    
    	// Make a new bucket called mymusic.
	bucketName,err := beego.AppConfig.String("bucketName")
	if err != nil {
        return err
    }  
	// location := "us-east-1"
	location := ""

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			fmt.Println("We already own bucket", bucketName)			
		} else {
			return err
		}
	} else {
		fmt.Println("Successfully created ", bucketName)
	}

	// Upload the zip file
	// objectName := "development/test2.jpeg"
	// filePath := "./images/test1.jpeg"
	objectName := file_s3
	filePath := file_src


	// contentType := "application/pdf"
	contentType := ""

	// Upload the zip file with FPutObject
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return err
	}

	fmt.Println("Successfully uploaded",objectName," of size ",info.Size)
	return nil
}

func GetS3(file string) (*minio.Object,error){
	bucketName,err := beego.AppConfig.String("bucketName")
    if err != nil {
        return nil,err
    }  

    minioClient,err := ConnS3Storage()
    if err != nil {
        return nil,err
    }    
    // fmt.Println(file)
    // object, err := minioClient.GetObject(context.Background(), "mekdi", "development/test2.jpeg", minio.GetObjectOptions{})
    object, err := minioClient.GetObject(context.Background(), bucketName , file, minio.GetObjectOptions{})
    if err != nil {
        return nil,err
    }

    return object,nil
}


func MapB64SaveFile(data_arr map[string]string,path string)(map[string]string,error) {

    var fileName = make(map[string]string)
    var err error    
    for key, data := range data_arr {	             
        if (data != ""){             
            fileName[key],err = B64SaveFile(data,path)                            
            if err != nil {
                break
                return fileName,err
            }  
        }
    }
    return fileName,err
}

func B64SaveFile(data string,pathnya string)(string,error) {

	onlyfile := data[strings.IndexByte(data, ',')+1:]
	mediatype := ""	
	switch data[0:strings.IndexByte(data,';')+1]{
		case "data:application/pdf;":
			mediatype=".pdf"
		case "data:image/png;":
			mediatype=".png"
		case "data:image/jpeg;":
			mediatype=".jpg"		
		case "data:image/jpg;":
			mediatype=".jpg"					
		default:
			mediatype="error"	
	}
	if mediatype == "error"{
		return "",errors.New("file tidak di izinkan untuk di upload")
	}	

    err := os.MkdirAll(pathnya, os.ModePerm)
    if err != nil {
        return "",err
    }

    fileNameBase := guid.New().String()+guid.New().String()    	

    reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(onlyfile))	    
    buff := bytes.Buffer{}
    _,err = buff.ReadFrom(reader)
    if err != nil {
        return "",err
    }
	
    fileName := pathnya+fileNameBase + mediatype

    err = ioutil.WriteFile(fileName, buff.Bytes(), 0600)
    if err != nil {
        return "",err
    }


    err = PutS3(fileName,fileNameBase + mediatype)
    if err != nil {
        return "",err
    }    

    err = os.Remove(fileName)
    if err != nil {
        return "",err
    } 

	return fileName,nil
}