package myrouter

import (
	"net/http"
	"strings"
)

type Tree struct {
	root *Node
}

func (t *Tree) add(path string, handler http.HandlerFunc) {
	// init route
	if t.root == nil {
		t.root = &Node{
			Path:    "",
			Part:    "",
			Handler: nil,
		}
	}

	parts := splitPath(path)
	n := t.root
	for _, part := range parts {
		if part == "" {
			panic("An empty segment is included in this path.")
		}
		child := n.MatchChild(part)
		if child == nil {
			child = &Node{
				Path:    path,
				Part:    part,
				Handler: handler,
				IsWild:  part[0] == ':' || part[0] == '*' || part[0] == '{',
			}
			n.Children = append(n.Children, child)
		}
		n = child
	}
	// n.Route = route
}

// get gets Node matched with path
func (t *Tree) get(path string) *Node {
	parts := splitPath(path)
	n := t.root
	for _, part := range parts {
		if n = n.MatchChild(part); n == nil {
			return nil
		}
	}

	return n
}

func splitPath(path string) []string {
	parts := strings.Split(path, "/")[1:]
	lastIndex := len(parts) - 1

	// remove a space after trailing slash.
	// ex) /test/ => " " + "test" + " "
	if parts[lastIndex] == "" {
		parts = parts[:lastIndex-1]
	}
	return parts
}
