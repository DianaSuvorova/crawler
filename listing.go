package main

import (
  "github.com/jinzhu/gorm"
  "github.com/PuerkitoBio/goquery"
)

type listingPage struct {
  doc *goquery.Document
  *listing
}

type listing struct {
  url string
  shopLink string
}

type shopSource struct {
  gorm.Model
  Url string `gorm:"unique_index"`
}

func newListingPage(url string) *listingPage {
  lp := new(listingPage)
  lp.listing = new(listing)
  lp.listing.url = url

  return lp
}

func (lp *listingPage) fetch() {
  done := make(chan bool);
  go func () {
    lp.doc, _ = goquery.NewDocument(lp.listing.url)
    done <- true
  }();
  for {
    select {
      case  <- done:
        return;
    }
  }
}

func (lp *listingPage)process() (zero []processor) {
  lp.fetch();
  first := lp.doc.Find("div.shop-name a").First()
  shop := new(shopSource);
  href, _ := first.Attr("href")
  shop.Url = href
  // db.CreateTable(&shopSource{})
  db.Create(shop)
  return
}

func (lp *listingPage)url() string {
	return lp.listing.url
}
