package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main(){
    fileName := flag.String("csv", "default.csv", "A csv file name to store your scraped data")
    flag.Parse()
    file, err := os.Create(*fileName)
    if err != nil{
        log.Fatalf("Cannot create file %q: %s \n", *fileName, err)
        return
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    //Write a CSV header
    writer.Write([]string{"Ring Name", "Real Name", "Notes"})


// Instantiate default collector
    fmt.Println("Hello world")

    c := colly.NewCollector()

    c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

    c.OnResponse(func(r *colly.Response) {
        fmt.Println("Visited", r.Request.URL)
    })

    c.OnHTML("table.wikitable", func(h *colly.HTMLElement) {
        //Print the link
        h.ForEach("tr", func(i int, el *colly.HTMLElement) {
            writer.Write([]string{
                el.ChildText("td:nth-child(1)") ,
                el.ChildText("td:nth-child(2)"),
                el.ChildText("td:nth-child(3)"),
            })
        })
        fmt.Println("Completed ")
    })


    c.Visit("https://en.wikipedia.org/wiki/List_of_All_Elite_Wrestling_personnel")
}
