package sitemap

import (
    "encoding/xml"
    "os"

    "github.com/go-xmlfmt/xmlfmt"
    "gophercise/sitemapGen/utils"
)

type site struct {
    Loc string `xml:"loc"`
}

type docShape struct {
    XMLName xml.Name `xml:"urlset"`
    Xmlns   string   `xml:"xmlns,attr"`
    URL     []site   `xml:"url"`
}

func WriteFile() {
    st := GetSiteMap()
    var sl []site
    for s, v := range st.Visited {
        if v {
            sl = append(sl, site{s})
        }
    }
    doc := docShape{xml.Name{}, "http://www.sitemaps.org/schemas/sitemap/0.9", sl}
    bt, err := xml.Marshal(doc)
    utils.CheckError(err, true)
    xmlFile, err := os.OpenFile("sitemap.xml", os.O_RDWR|os.O_CREATE, 0660)
    utils.CheckError(err, true)
    s := xmlfmt.FormatXML(string(bt), "", "    ")
    _, err = xmlFile.Write([]byte(s))
    utils.CheckError(err, true)
}
