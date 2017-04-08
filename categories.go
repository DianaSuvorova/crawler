package main

import (
  "fmt"
)

type CategoryPage struct {
  *Page
  *Category
}

type Category struct {
  title string
  items string
}

func NewCategoryPage(url string) *CategoryPage {
  fmt.Println(url)
  cp := new(CategoryPage)
  cp.Page = NewPage(url)

  return cp
}

func (cp *CategoryPage)Process() (zero []Processor) {
  fmt.Println("CategoryPage", cp.Page.Url)
  return
}

func (cp *CategoryPage)Url() string {
	return cp.Page.Url
}
