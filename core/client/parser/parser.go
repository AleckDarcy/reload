package parser

import (
	"bytes"
	"io"

	"golang.org/x/net/html"
)

func GetElementByClass(node *html.Node, class string) *html.Node {
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		for _, attr := range n.Attr {
			if attr.Key == "class" && attr.Val == class {
				return n
			}
		}

		if n := GetElementByClass(n, class); n != nil {
			return n
		}
	}

	return nil
}

func GetJSON(node *html.Node) string {
	buf := &bytes.Buffer{}
	w := io.Writer(buf)
	html.Render(w, node.FirstChild)

	return html.UnescapeString(buf.String())
}
