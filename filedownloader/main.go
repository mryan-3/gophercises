package main

import (
    "fmt"
    "io"
    "log"
    "net/http"
)


func main(){
    fmt.Println("THello from the moon")

    res, err := http.Get("http://www.google.com/robots.txt")
    if err != nil {
        log.Fatal(err)
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()

    if res.StatusCode > 299 {
        log.Fatal("Response failed with status code %d and \nbody: %s\n", res.StatusCode, body)
    }

    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("%s", body)
}
