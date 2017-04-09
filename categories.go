package main

import (
  "fmt"
  "github.com/PuerkitoBio/goquery"
  "regexp"
)

type CategoryPage struct {
  Doc *goquery.Document
  url string
  *Category
}

type Category struct {
  title string
  items string
}

func NewCategoryPage(url string) *CategoryPage {
  cp := new(CategoryPage)
  cp.url = url
  cp.Doc, _ = goquery.NewDocument(url)

  return cp
}

func (cp *CategoryPage)Process() (zero []Processor) {
  fmt.Println(cp.url)
  cp.Doc.Find("div.mt-xs-2 span").Each(func(i int, s *goquery.Selection) {
    match, _ := regexp.MatchString("\\([,0-9]* items\\)", s.Text())
     if (match) {
       fmt.Println(s.Text())
     }
  })
  return
}

func (cp *CategoryPage)Url() string {
	return cp.url
}
