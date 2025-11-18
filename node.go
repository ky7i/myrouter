package myrouter

import (
	"net/http"
)

// type Route struct {
// 	Path    string `json:"path"`
// 	Handler http.HandlerFunc
// }

type Node struct {
	Path     string           `json:"path"`
	Part     string           `json:"part"`
	Children []*Node          `json:"children`
	IsWild   bool             `json:"isWild`
	Handler  http.HandlerFunc `json:"Handler"`
}

// func (n *Node) Insert(pattern string, route Route) {
// 	parts := strings.Split(pattern, "/")[1:]
//
// 	for _, part := range parts {
// 		child := n.MatchChild(part)
// 		if child == nil {
// 			child = &Node{
// 				Part:   part,
// 				IsWild: part[0] == ':' || part[0] == '*' || part[0] == '{',
// 			}
// 			n.Children = append(n.Children, child)
// 		}
// 		n = child
// 	}
//
// 	n.Route = route
// }
//
// func (n *Node) Search(path string) Route {
// 	parts := strings.Split(path, "/")[1:]
//
// 	for _, part := range parts {
// 		child := n.MatchChild(part)
// 		if child == nil {
// 			return Route{}
// 		}
// 		n = child
// 	}
//
// 	return n.Route
// }

func (n *Node) MatchChild(part string) *Node {
	for i := range n.Children {
		if n.Children[i].Part == part || n.Children[i].IsWild {
			return n.Children[i]
		}
	}
	return nil
}
