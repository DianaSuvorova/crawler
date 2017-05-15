package main

import (
	"encoding/xml"
)

type siteMapMetaPage struct {
 *page
 *siteMapMeta
}

type siteMapMeta struct {
 Links []xmlLink `xml:"sitemap"`
}

func newSiteMapMetaPage(url string) *siteMapMetaPage {
  smmp := new(siteMapMetaPage)
	smmp.page = newPage(url)
  return smmp
}

func (smmp *siteMapMetaPage) process() (siteMapPages []processor) {
	success := smmp.page.fetch();
	if success {
		xml.Unmarshal([]byte(smmp.page.body), &smmp.siteMapMeta)
		for _, link := range smmp.siteMapMeta.Links {
			siteMap := newSiteMapPage(link.String())
			siteMapPages = append(siteMapPages, siteMap)
		}
	}
	return
}

func (smmp *siteMapMetaPage)url() string {
	return smmp.page.url
}
