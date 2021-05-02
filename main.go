package main

import (
	"os"
	"strings"

	"github.com/PurdueCho/jobScrapper/scrapper"
	"github.com/labstack/echo/v4"
)

func handleHome(c echo.Context) error {
	return c.File("index.html")
}

const fileNAME string = "jobs.csv"

func handleScrape(c echo.Context) error {
	defer os.Remove(fileNAME)
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scarpe(term)
	return c.Attachment(fileNAME, fileNAME)
}

func main() {
	//scrapper.Scarpe("term")

	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}
