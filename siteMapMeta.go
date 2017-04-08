package main

import (
	"encoding/xml"
)

type SiteMapMetaPage struct {
 *Page
 *SiteMapMeta
}

type SiteMapMeta struct {
 Links []XmlLink `xml:"sitemap"`
}

func NewSiteMapMetaPage(url string) *SiteMapMetaPage {
  smmp := new(SiteMapMetaPage)
  smmp.Page = NewPage(url)

  xml.Unmarshal([]byte(smmp.Page.Body), &smmp.SiteMapMeta)
  return smmp
}

func (smmp *SiteMapMetaPage)Process() (siteMapPages []Processor) {
  for _, link := range smmp.SiteMapMeta.Links {
    siteMap := NewSiteMapPage(link.String())
    siteMapPages = append(siteMapPages, siteMap)
  }
  return
}

func (smmp *SiteMapMetaPage)Url() string {
	return smmp.Page.Url
}
