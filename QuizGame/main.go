package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main (){
    csvFile := flag.String("csv", "problems.csv", "A csv file in the format question, answer")
    flag.Parse()
    file, err := os.Open(*csvFile)

    if err != nil {
        log.Fatal("Error while opening the file ", *csvFile)
    }

    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()

    if err != nil{
        fmt.Println("Error while reading the records")
    }
    parseRecords(records)
}

func parseRecords(records [][]string) []problem{
    ret := make([]problem, len(records))
    for i, record := range records{
        ret[i] = problem{
            q: record[0],
            a: record[1],
        }
    }
    fmt.Println(ret)
    return ret
}

type problem struct{
    q string
    a string
}
