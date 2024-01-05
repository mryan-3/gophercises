package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)

func main (){
    file, err := os.Open("./problems.csv")

    if err != nil {
        log.Fatal("Error while opening the file")
    }

    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()

    if err != nil{
        fmt.Println("Error while reading the records")
    }

    for _, eachRecord := range records{
        fmt.Println(eachRecord)
    }
}
