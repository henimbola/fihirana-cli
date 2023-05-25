package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

var version = "0.0.1-alpha"
var author = "ANDRIANATOANDRO Tsihala Henimbola"

func main() {
	// Load the HTML document
	number := flag.Int("n", 0, "an argument")
	shortVersion := flag.Bool("v", false, "show the version and exit")
	showVersion := flag.Bool("version", false, "show the version and exit")
	showNeoFetch := flag.Bool("info", false, "show a styled text")

	flag.Parse()

	if *showNeoFetch {
		fmt.Printf("Version : %v \n", version)
		fmt.Printf("Author : %v \n", version)
		fmt.Println("Release Date: don't know yet")
	}

	if *showVersion || *shortVersion {
		fmt.Println(version)
	}

	if *number != 0 {
		getLyrics(*number)
	}
}

func getLyrics(number int) {
	doc, err := goquery.NewDocument(fmt.Sprintf("https://fihirana.org//par-numero/recherche-multiple?fihirana_numbers=%d", number))
	if err != nil {
		log.Fatal(err)
	}

	// Find the title element and print its text
	link, _ := doc.Find(".entry-content h4 > a").Attr("href")

	doc2, _ := goquery.NewDocument(link)

	fmt.Printf("Lohateny : %v \n\n", doc2.Find(".entry-title").Text())
	doc2.Find(".entry-content > p").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		fmt.Println()
	})
}
