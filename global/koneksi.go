package global
import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"errors"
	_ "github.com/denisenkom/go-mssqldb"	
	"github.com/funujikai/godb"
	"github.com/funujikai/godb/adapters/mssql"	
	beego "github.com/beego/beego/v2/server/web"


	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)


// Connection URI
// const uri = "mongodb://10.61.3.83:27017/?maxPoolSize=20&w=majority"
const uri = "mongodb://10.61.3.83:27017/?readPreference=primary&directConnection=true&ssl=false&replicaSet=rs0"
func Conn_mongodb()(*mongo.Client,error) {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return client,err
	}
	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return client,err
	}
	fmt.Println("Successfully connected and pinged.")

	return client,nil
}




func ConnINTEGRASI_MEKAAR() (*godb.DB,error) {
	SQL_STRING,_ := beego.AppConfig.String("sqlconnMEKAAR_INTEGRASI")

	conn, err := godb.Open(mssql.Adapter, SQL_STRING)
	if err != nil {
		return nil,err
	}

	err = conn.Ping()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}


	return conn,nil
}



func ConnSSO() (*godb.DB,error) {
	SQL_STRING,_ := beego.AppConfig.String("sqlconnSSO")

	conn, err := godb.Open(mssql.Adapter, SQL_STRING)
	if err != nil {
		return nil,err
	}

	err = conn.Ping()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}


	return conn,nil
}


func Conn() (*godb.DB,error) {
	SQL_STRING,_ := beego.AppConfig.String("sqlconnPMM")

	conn, err := godb.Open(mssql.Adapter, SQL_STRING)
	if err != nil {
		return nil,err
	}

	err = conn.Ping()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}


	return conn,nil
}



func Conn15() (*godb.DB,error) {
	SQL_STRING,_ := beego.AppConfig.String("sqlconn15")

	conn, err := godb.Open(mssql.Adapter, SQL_STRING)
	if err != nil {
		return nil,err
	}

	err = conn.Ping()
	if err != nil {
		return nil,errors.New("KONEKSI KE DATABASE PENUH, COBA BEBERAPA MENIT LAGI")
	}


	return conn,nil
}



func ConnS3Storage()(*minio.Client,error){
    
	endpoint := "pnmdc-cluster-cohesity.pnm.co.id:3000"
	accessKeyID := "i4Xa_j1mUZE0xtvWZzdCGLIuKGoTGBE2mZBt0dWmcKM"
	secretAccessKey := "SlZTrEhAEH3ZUwOnkzmWPEMQAj89bqjpXUt_PAgsPDI"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return minioClient,err
	}

	return minioClient,nil
}