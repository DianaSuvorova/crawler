package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "regexp"
  "strings"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type categoryPage struct {
  doc *goquery.Document
  *category
}

type category struct {
  gorm.Model
  url string
  Items string
}

func NewCategoryPage(url string) *categoryPage {
  cp := new(categoryPage)
  cp.category = new(category)
  cp.category.url = url
  cp.Doc, _ = goquery.NewDocument(url)

  return cp
}

func (cp *categoryPage)process() (zero []processor) {
  cp.Doc.Find("div.mt-xs-2 span").Each(func(i int, s *goquery.Selection) {
    match, _ := regexp.MatchString("\\([,0-9]* items\\)", s.Text())
     if (match) {
       result := strings.TrimSpace(s.Text())
       cp.category.Items = result
       cp.category.write()
     }
  })
  return
}

func (cp *categoryPage)url() string {
	return cp.category.url
}

func(c *category)write() {
  db, err := gorm.Open("mysql", "-connection tbd")
  db.LogMode(true)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()
  fmt.Println(c)
  db.CreateTable(&category{})

  // db.Create(c)
}
