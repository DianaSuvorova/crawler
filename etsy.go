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
	connection := flag.Args()[0]
	db, err = gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	entryUrl := "https://www.etsy.com/robots.txt"
	entryPage := newRobotsPage(entryUrl)

	startQueuer(entryPage)
}
