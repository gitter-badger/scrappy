package scrappy

import (
	"golang.org/x/net/html"
)

// All, group of methods that return all occurrence
type All struct {
	*Scrappy
}

// Depth return nodes using first depth algorithm
func (a *All) Depth(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	if Validate(node, filters...) {
		result = append(result, node)
	}
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		result = append(result, a.Depth(node, filters...)...)
	}
	return result
}

// Breadth return nodes using first breadth algorithm
func (a *All) Breadth(node *html.Node, filters ...FilterFunc) []*html.Node {
	var breadth func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node
	breadth = func(nodes []*html.Node, result []*html.Node, filters ...FilterFunc) []*html.Node {
		var next []*html.Node
		for _, elm := range nodes {
			for node := elm.FirstChild; node != nil; node = node.NextSibling {
				if Validate(node, filters...) {
					result = append(result, node)
				}
				next = append(next, node)
			}
		}
		if len(next) > 0 {
			return breadth(next, result, filters...)
		}
		return result
	}
	return breadth([]*html.Node{node}, []*html.Node{}, filters...)
}

// Parent return parent nodes that matches with given filters
func (a *All) Parent(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.Parent; node != nil; node = node.Parent {
		if Validate(node, filters...) {
			result = append(result, node)
		}
	}
	return result
}

// Child return child nodes that matches with given filters
func (a *All) Child(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.FirstChild; node != nil; node = node.NextSibling {
		if Validate(node, filters...) {
			result = append(result,node)
		}
	}
return result
}

// Next return next sibling nodes that matches with given filters
func (a *All) NextSibling(root *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := root.NextSibling; node != nil; node = node.NextSibling {
		if node.LastChild != nil && node.PrevSibling.Data != root.Data && node.Parent != root {
			if Validate(node, filters...) {
				result = append(result, node)
			}
		}
	}
	return result
}

// Prev return prev sibling nodes that matches with given filters
func (a *All) PrevSibling(node *html.Node, filters ...FilterFunc) []*html.Node {
	var result []*html.Node
	for node := node.PrevSibling; node != nil; node = node.PrevSibling {
		if Validate(node, filters...) {
			result = append(result, node)
		}
	}
	return result
}