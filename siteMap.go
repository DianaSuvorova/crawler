package main

import (
	"encoding/xml"
)

type SiteMapPage struct {
 *Page
 *SiteMap
}

type SiteMap struct {
 Links []XmlLink `xml:"url"`
}

func NewSiteMapPage(url string) *SiteMapPage {
  smp := new(SiteMapPage)
  smp.Page = NewPage(url);


  xml.Unmarshal([]byte(smp.Page.Body), &smp.SiteMap)
  return smp
}

func (smp *SiteMapPage)Process() (pages []Processor) {
	if (len(smp.Links) > 0) {
		page := NewCategoryPage(smp.Links[0].String())
		pages = append(pages, page)
	}
  return
}

func (smp *SiteMapPage)Url() string {
	return smp.Page.Url
}
