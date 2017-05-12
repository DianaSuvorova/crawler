package main

import (
	"encoding/xml"
)

type siteMapPage struct {
 *page
 *siteMap
}

type siteMap struct {
 Links []xmlLink `xml:"url"`
}

func newSiteMapPage(url string) *siteMapPage {
  smp := new(siteMapPage)
  smp.page = newPage(url)

  return smp
}

func (smp *siteMapPage)process() (pages []processor) {
	smp.page.fetch();
	xml.Unmarshal([]byte(smp.page.body), &smp.siteMap)

	if (len(smp.Links) > 0) {
		link := smp.Links[0].String()
		page := newListingPage(link)
		// page := newCategoryPage(link)
		pages = append(pages, page)
	}
  return
}

func (smp *siteMapPage)url() string {
	return smp.page.url
}
