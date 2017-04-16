package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "regexp"
  "github.com/jinzhu/gorm"
  "strconv"
  "strings"
)

type categoryPage struct {
  doc *goquery.Document
  *category
}

type category struct {
  gorm.Model
  Url string
  Items int
}

func NewCategoryPage(url string) *categoryPage {
  cp := new(categoryPage)
  cp.category = new(category)
  cp.category.Url = url
  cp.doc, _ = goquery.NewDocument(url)

  return cp
}

func (cp *categoryPage)process() (zero []processor) {
  cp.doc.Find("div.mt-xs-2 span").Each(func(i int, s *goquery.Selection) {
    re, err := regexp.Compile("\\(([,0-9]+) items\\)")
    if (err != nil) {
      fmt.Println(err)
    }
    res := re.FindAllStringSubmatch(s.Text(), -1)
    if (res != nil ) {
      cp.category.Items, _ = strconv.Atoi(strings.Replace(res[0][1], ",","", -1))
      cp.category.write()
    }
  })
  return
}

func (cp *categoryPage)url() string {
	return cp.category.Url
}

func(c *category)write() {
  //db.CreateTable(&category{})
  fmt.Println(c)
  db.Create(c)
}
