package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"flag"
)

var db *gorm.DB


func main() {
	var err error
	flag.Parse()
	connection := "boris:B@ckspace123@tcp(54.215.211.253:3306)/etsy?parseTime=true"
	//connection := flag.Args()[0]
	db, err = gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// entryUrl := "https://www.etsy.com/robots.txt"
	// entryPage := newRobotsPage(entryUrl)

	entryPage := newSiteMapMetaPage("https://www.etsy.com/dynamic-sitemaps.xml?sitemap=shoplisting_index2&min=5040064&max=5050063")
	startQueuer(entryPage)
}
