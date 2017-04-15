package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "regexp"
  "strings"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type CategoryPage struct {
  Doc *goquery.Document
  *Category
}

type Category struct {
  gorm.Model
  Url string
  Items string
}

func NewCategoryPage(url string) *CategoryPage {
  cp := new(CategoryPage)
  cp.Category = new(Category)
  cp.Category.Url = url
  cp.Doc, _ = goquery.NewDocument(url)

  return cp
}

func (cp *CategoryPage)Process() (zero []Processor) {
  cp.Doc.Find("div.mt-xs-2 span").Each(func(i int, s *goquery.Selection) {
    match, _ := regexp.MatchString("\\([,0-9]* items\\)", s.Text())
     if (match) {
       result := strings.TrimSpace(s.Text())
       cp.Category.Items = result
       cp.Category.write()
     }
  })
  return
}

func (cp *CategoryPage)Url() string {
	return cp.Category.Url
}

func(c *Category)write() {
  db, err := gorm.Open("mysql", "--connection tbd")
  db.LogMode(true)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  db.CreateTable(&Category{})

  db.Create(c)
}
