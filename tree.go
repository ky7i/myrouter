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
		t.root = &Node{}
	}

	parts := strings.Split(path, "/")[1:]
	n := t.root
	for _, part := range parts {
		if part == "" {
			panic("An empty segment or a trailing slash is included.")
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
	parts := strings.Split(path, "/")[1:]
	n := t.root
	for _, part := range parts {
		if n = n.MatchChild(part); n == nil {
			return nil
		}
	}

	return n
}
