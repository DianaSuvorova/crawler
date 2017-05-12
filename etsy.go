package main


import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"flag"
	"runtime"
)

var db *gorm.DB


func main() {
	var err error
	flag.Parse()
	connection := "boris:B@ckspace123@tcp(54.215.211.253:3306)/etsy"
	//flag.Args()[0]
	db, err = gorm.Open("mysql", connection)
	print("goroutines: ", runtime.NumGoroutine())
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	entryUrl := "https://www.etsy.com/robots.txt"
	entryPage := newRobotsPage(entryUrl)

	startQueuer(entryPage)
	print("goroutines: ", runtime.NumGoroutine())
}
