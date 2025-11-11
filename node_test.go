package myrouter

import (
	"net/http"
	"testing"
)

func TestMatchChildMatched(t *testing.T) {
	word := ""
	handler1 := func(w http.ResponseWriter, r *http.Request) { word = "handler1" }
	handler2 := func(w http.ResponseWriter, r *http.Request) { word = "handler2" }

	routes := []Route{
		{Path: "/test1", Handler: handler1},
		{Path: "/test2", Handler: handler2},
	}
	tree := Node{}
	for _, route := range routes {
		tree.Insert(route.Path, route)
	}
	path := "/test1"
	child := tree.MatchChild(path)
	if child == nil || child.Route.Path != "/test1" {
		t.Errorf("Path: %q, Handler: %q got, want Path: '/test1', Handler: 'Handler1'", child.Route.Path, child.Route.Handler)
	}
	child.Route.Handler(nil, nil)
	if word != "handler1" {
		t.Errorf("Expected handler is Handler1, got %q", word)
	}
}
