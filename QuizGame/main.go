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
    problems := parseRecords(records)
    for i, p := range problems{
        fmt.Printf("Problem %d: %s = \n", i+1, p.q)
        var answer string
        fmt.Scanf("%s\n", &answer)
        if answer == p.a {
            fmt.Println("Correct")
        } else {
            fmt.Println("Incorrect")
        }
    }
}

func parseRecords(records [][]string) []problem{
    ret := make([]problem, len(records))
    for i, record := range records{
        ret[i] = problem{
            q: record[0],
            a: record[1],
        }
    }
    return ret
}

type problem struct{
    q string
    a string
}
