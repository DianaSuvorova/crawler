package main

import (
  "github.com/jinzhu/gorm"
  "github.com/PuerkitoBio/goquery"
  "regexp"
  "strconv"
  "strings"
)

type shopPage struct {
  doc *goquery.Document
  *shopSource
  *shopRecord
}

type shopRecord struct {
  gorm.Model
  TotalNumSold int
  Url string
  RunId uint
}

func newShopPage(source shopSource, runId uint) *shopPage  {
  sp := new(shopPage)
  sp.shopSource = &source
  sp.shopRecord = new(shopRecord)
  sp.shopRecord.Url = source.Url
  sp.shopRecord.RunId = runId

  return sp;
}

func (sp * shopPage) process() {
  var err error;
  sp.doc, err = goquery.NewDocument(sp.shopRecord.Url)
  if (err == nil) {
    info := sp.doc.Find(".show-lg .trust-signal-row")
    if (info.Length() == 0) {
      sp.shopSource.Deleted = true;
      db.Save(sp.shopSource)
    } else {
      sold := info.Find(".mr-xs-2.pr-xs-2.br-xs-1").Last().Text()
      re := regexp.MustCompile("([0-9]+)")
      totalNumSold, _ := strconv.Atoi(re.FindAllString(sold, -1)[0])
      sp.shopRecord.TotalNumSold = totalNumSold
      joined := strings.Replace(info.Find(".etsy-since").Text(), "On Etsy since", "", -1)
      sp.shopSource.Joined = joined;
      //db.CreateTable(&shopRecord{})
      db.Create(sp.shopRecord)
    }
  }
}
