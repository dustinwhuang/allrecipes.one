package main

import (
	"golang.org/x/net/html"
	"io"
)

func dfs(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" {
		return n.FirstChild.Data
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title := dfs(c)
		if title != "" {
			return title
		}
	}

	return ""
}

func getTitle(r io.Reader) string {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	return dfs(doc)
}
