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
	connection := "boris:B@ckspace123@tcp(54.215.211.253:3306)/etsy"
	//flag.Args()[0]
	db, err = gorm.Open("mysql", connection)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	entryUrl := "https://www.etsy.com/robots.txt"
	entryPage := newRobotsPage(entryUrl)

	// entryUrl := "https://www.etsy.com/listing/521442201/maui-sunrise-2"
	// entryPage := newListingPage(entryUrl)

	startQueuer(entryPage)
}
