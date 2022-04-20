package global

import (
    // beego "github.com/beego/beego/v2/server/web"
	"fmt"
    "os"
	"io/ioutil"
    // "time"
    "github.com/alexmullins/zip"
    "path/filepath"
)

func ZipWriter(baseFolder string,output string)error {    
    baseFolderClean := filepath.Clean(baseFolder)
    // current_time := time.Now()
    // Get a Buffer to Write To
    // outFile, err := os.Create(output+"_"+current_time.Format("20060102150405")+".zip")
    outFile, err := os.Create(output+".zip")
    if err != nil {
        return err
    }
    // defer outFile.Close()

    // Create a new zip archive.
    zipw := zip.NewWriter(outFile)    

    // Add some files to the archive.    
    addFiles(zipw, baseFolderClean, "")


    // Make sure to check the error on Close.
    err = zipw.Close()
    if err != nil {
        fmt.Println(err)
    }    

    err = outFile.Close()
    if err != nil {
        fmt.Println(err)
    }

    return nil
}


func addFiles(zipw *zip.Writer, basePath, baseInZip string) {


    // Open the Directory
    files, err := ioutil.ReadDir(basePath)
    if err != nil {
        fmt.Println(err)
    }

    for _, file := range files {
        if !file.IsDir() {
            dat, err := ioutil.ReadFile(filepath.Clean(basePath+"/"+file.Name()))
            if err != nil {
                fmt.Println("error 1",err)
            }

			// passwordfile,err := beego.AppConfig.String("PasswordFile")
			// if err != nil {
            //     fmt.Println("error 1",err)
            // }

            // Add some files to the archive.
            f, err := zipw.Create(baseInZip+file.Name())
            if err != nil {
                fmt.Println("error 2",err)
            }
            _, err = f.Write(dat)
            if err != nil {
                fmt.Println("error 3",err)
            }
        } else if file.IsDir() {

            // Recurse
            newBase := basePath + "/" + file.Name() + "/"
            fmt.Println("Recursing and Adding SubDir: " + file.Name())
            fmt.Println("Recursing and Adding SubDir: " + newBase)

            addFiles(zipw, newBase, baseInZip + file.Name() + "/")
        }        
    }
}