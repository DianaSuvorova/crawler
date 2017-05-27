package main


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"flag"
	"sync"
)

var db *gorm.DB
var wg sync.WaitGroup


func main() {
	var err error
	flag.Parse()
	flags := flag.Args()
	connection := flags[0]
	option := flags[1]


	db, err = gorm.Open("mysql", connection)
	db.LogMode(true)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	if (option == "parser") {
		runId := newRunLog().Id();
		startParsing(runId);
	} else if (option == "crawler") {
		entryUrl := "https://www.etsy.com/robots.txt"
		entryPage := newRobotsPage(entryUrl)
		startQueuer(entryPage)
	} else {
		print("Second arg should be parser or crawler");
	}

}
