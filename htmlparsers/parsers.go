package htmlparsers

import (
	"strings"

	"golang.org/x/net/html"
)

func GetValFromNode(n *html.Node, data, key string) string {
	if n.Type == html.ElementNode && n.Data == data {
		for _, attr := range n.Attr {
			if attr.Key == key {
				return attr.Val
			}
		}
	}
	return ""
}

func DeepGetValsFromNode(n *html.Node, data, key string) []string {
	srcs := []string{}

	if n.Type == html.ElementNode && n.Data == data {
		for _, attr := range n.Attr {
			if attr.Key == key {
				srcs = append(srcs, attr.Val)
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		srcs = append(srcs, DeepGetValsFromNode(c, data, key)...)
	}

	return srcs
}

func FindAllImgSrcVals(htmlStr string) []string {
	rootNode, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return []string{}
	}
	return DeepGetValsFromNode(rootNode, "img", "src")
}
