package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func downloadFile( urlPath string)(err error){
    //Build filename from the url path
    fileUrl, err := url.Parse(urlPath)
    if err != nil {
        log.Fatal(err)
    }
    path := fileUrl.Path
    segments := strings.Split(path, "/")
    fileName := segments[len(segments)-1]

    //Create the downloadFile
    file, err := os.Create(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    //Get the data
    resp, err := http.Get(urlPath)
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    //Check the status
    if resp.StatusCode != http.StatusOK{
        log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, resp.Body)
    }

    //Write the body to the file
    size, err := io.Copy(file, resp.Body)
    if err != nil{
        log.Fatal(err)
    }
    defer file.Close()
    fmt.Printf("Downloaded a file %s with size %d", fileName, size)

    return nil
}


func main(){
   fmt.Println("THello from the moon")
    downloadFile("https://www.google.com/robots.txt")

}
