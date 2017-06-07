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
  Joined string
  Deleted bool
}

func newListingPage(url string) *listingPage {
  lp := new(listingPage)
  lp.listing = new(listing)
  lp.listing.url = url

  return lp
}

func (lp *listingPage) process(availSpaceInQueue int) (zero []processor) {
  var err error;
  lp.doc, err = goquery.NewDocument(lp.listing.url)
  if (err == nil) {
    first := lp.doc.Find("div.shop-name a").First()
    shop := new(shopSource);
    href, _ := first.Attr("href")
    shop.Url = href
    // db.CreateTable(&shopSource{})
    db.Create(shop)
  }
  return
}

func (lp *listingPage)url() string {
	return lp.listing.url
}
