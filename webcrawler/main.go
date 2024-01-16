package main

import (
    "fmt"
    "github.com/gocolly/colly"
)

func main(){
// Instantiate default collector
    fmt.Println("Hello world")

    c := colly.NewCollector()

    c.OnHTML("*", func(h *colly.HTMLElement) {
        //Print the link
        fmt.Printf(" found: %q -> %s\n", h.Text)
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("visiting", r.URL.String())
    })
        c.OnResponse(func(r *colly.Response) {
        fmt.Println("Visited", r.Request.URL)
    })

    c.Visit("https://hackerspaces.org/")
}
