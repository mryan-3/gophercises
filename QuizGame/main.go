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
        fmt.Printf("Problem %d: %s = ", i+1, p.q)
        answerCh := make(chan string)
        go func(){
            var answer string
            fmt.Scanf("%s\n", &answer)
            answerCh <- answer
        } ()

        select{
        case <- timer.C:
            fmt.Printf("You scored %d out of %d. \n", correct, len(problems))
            return
        case answer := <-answerCh:
            if(answer == p.a) {
                correct++
            }
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
