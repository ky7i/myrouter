package myrouter

import (
	"log"
	"net/http"
	"strings"
)

type Tree struct {
	route *Node
}

func (t *Tree) add(path string, handler func(http.ResponseWriter, *http.Request)) {
	// init route
	if t.route == nil {
		t.route = &Node{}
	}

	parts := strings.Split(path, "/")[1:]
	n := t.route
	for _, part := range parts {
		child := n.MatchChild(part)
		if child == nil {
			child = &Node{
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
	n := t.route
	for _, part := range parts {
		child := n.MatchChild(part)
		if child == nil {
			return &Node{}
		}
		n = child
	}

	return n
}
