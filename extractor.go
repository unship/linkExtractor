package linkExtractor

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)



func Extract(h string, u string) ([]string, error) {
	URL,err:=url.Parse(u)
	if err!=nil{
		return []string{},err
	}
	doc, err := html.Parse(strings.NewReader(h))
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", u, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
