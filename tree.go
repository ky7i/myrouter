package myrouter

import (
	"log"
	"net/http"
	"strings"
)

type Tree struct {
	root *Node
}

func (t *Tree) add(path string, handler func(http.ResponseWriter, *http.Request)) {
	// init route
	if t.root == nil {
		t.root = &Node{}
	}

	parts := strings.Split(path, "/")[1:]
	n := t.root
	for _, part := range parts {
		child := n.MatchChild(part)
		if child == nil {
			child = &Node{
				Path:    path,
				Part:    part,
				Handler: handler,
				IsWild:  part[0] == ':' || part[0] == '*' || part[0] == '{',
			}
			n.Children = append(n.Children, child)
			log.Printf("node (path: %q) is successfuly add.", path)
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
