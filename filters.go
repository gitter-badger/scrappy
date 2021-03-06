package scrappy

import (
	"golang.org/x/net/html"
	"strings"
)

// FilterFunc is the general definition of a node filter
type FilterFunc func(node *html.Node) bool

// Tag  s a filter func that return a node that matches with a given string
func Tag(val string) FilterFunc {
	return func(node *html.Node) bool { return node.Type == html.ElementNode && node.Data == val }
}

// Text is a filter func that return a node that matches with a given string
func Text(val string) FilterFunc {
	return func(node *html.Node) bool {
		return node.Type == html.TextNode && node.Data == val
	}
}

// Attr is a filter func that return a node that matches with a given string
func Attr(val string) FilterFunc {
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if a.Key == val {
				return true
			}
		}
		return false
	}
}

// Values is a filter func that return a node that matches with a given string
func Value(val string) FilterFunc {
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if a.Val == val {
				return true
			}
		}
		return false
	}
}

// AttrVal is a filter func that return a node that matches with a pair attr/value
func AttrVal(attr string, val string) FilterFunc {
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if a.Key == attr && a.Val == val {
				return true
			}
		}
		return false
	}
}

// ContainTag is a filter func that return a node with a tag that contain a given string
func ContainsTag(val string) FilterFunc {
	return func(node *html.Node) bool { return node.Type == html.ElementNode && strings.Contains(node.Data, val) }
}

// ContainText is a filter func that return a node that contain a given string
func ContainsText(val string) FilterFunc {
	return func(node *html.Node) bool {
		return node.Type == html.TextNode && strings.Contains(node.Data, val)
	}
}

// ContainAttr is a filter func that return a node with an attr that contain a given string
func ContainsAttr(val string) FilterFunc {
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if strings.Contains(a.Key, val) {
				return true
			}
		}
		return false
	}
}

// AttrValues is a filter func that return a node with an attr value that contain a given string
func ContainsValue(val string) FilterFunc {
	return func(node *html.Node) bool {
		for _, a := range node.Attr {
			if strings.Contains(a.Val, val) {
				return true
			}
		}
		return false
	}
}
