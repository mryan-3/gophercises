package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main (){
    csvFile := flag.String("csv", "problems.csv", "A csv file in the format question, answer")
    timeLimit := flag.Int("limit", 1, "The time limit for the quiz game in seconds")
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

    correct := 0
    problems := parseRecords(records)

    timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)


    for i, p := range problems{
        select {
        case <-timer.C:
            fmt.Printf("You scored %d  out of %d. \n", correct, len(problems))
            return
        default:
            fmt.Printf("Problem %d: %s = \n", i+1, p.q)
            var answer string
            fmt.Scanf("%s\n", &answer)
            if answer == p.a {
              correct += 1
            fmt.Printf("Problem %d: %s = \n", i+1, p.q)
            var answer string
            fmt.Scanf("%s\n", &answer)
            if answer == p.a {
                  correct += 1
            }
          }
          fmt.Printf("You scored %d out of %d \n", correct, len(problems))
          }
        }
}

func parseRecords(records [][]string) []problem{
    ret := make([]problem, len(records))
    for i, record := range records{
        ret[i] = problem{
            q: record[0],
            a: strings.TrimSpace(record[1]),
        }
    }
    return ret
}

type problem struct{
    q string
    a string
}
