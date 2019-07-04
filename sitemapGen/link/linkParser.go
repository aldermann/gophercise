package link

import (
    "io"
    "strings"

    "golang.org/x/net/html"
    "gophercise/sitemapGen/utils"
)

func Parse(r io.Reader) (LinksList, error) {

    var Links LinksList

    var traverse func(node *html.Node, lvl string) string

    traverse = func(node *html.Node, lvl string) string {
        if node.Type == html.TextNode {
            return strings.TrimSpace(node.Data)
        }
        var childDatum []string
        for c := node.FirstChild; c != nil; c = c.NextSibling {
            childData := traverse(c, lvl + "-")
            if childData != "" {
                childDatum = append(childDatum, childData)
            }
        }
        text := strings.Join(childDatum, " ")
        if node.Type == html.ElementNode && node.Data == "a" {
            var href string
            for _, v := range node.Attr {
                if v.Key == "href" {
                    href = v.Val
                }
            }
            Links = append(Links, Link{Text: text, Href: href})
        }
        return text
    }

    root, err := html.Parse(r)
    utils.CheckError(err, false)
    if err != nil {
        return nil, err
    }
    traverse(root, "")
    return Links, nil
}
