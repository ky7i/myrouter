package myrouter

import (
	"net/http"
)

type Node struct {
	Path     string           `json:"path"`
	Part     string           `json:"part"`
	Children []*Node          `json:"children`
	IsWild   bool             `json:"isWild`
	Handler  http.HandlerFunc `json:"Handler"`
}

func (n *Node) MatchChild(part string) *Node {
	for i := range n.Children {
		if n.Children[i].Part == part || n.Children[i].IsWild {
			return n.Children[i]
		}
	}
	return nil
}
