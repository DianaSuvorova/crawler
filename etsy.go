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
	connection := flag.Args()[0]

	db, err = gorm.Open("mysql", connection)
	db.LogMode(true)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	// entryUrl := "https://www.etsy.com/robots.txt"
	// entryPage := newRobotsPage(entryUrl)

	// entryUrl := "https://www.etsy.com/listing/521442201/maui-sunrise-2"
	// entryPage := newListingPage(entryUrl)

	//startQueuer(entryPage)

	startParsing();
}
