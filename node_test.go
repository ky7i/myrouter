package myrouter

import (
	"net/http"
	"testing"
)

func TestMatchChildMatched(t *testing.T) {
	word := ""
	handler1 := func(w http.ResponseWriter, r *http.Request) { word = "handler1" }
	handler2 := func(w http.ResponseWriter, r *http.Request) { word = "handler2" }

	routesList := [][]Route{{
		{Path: "/test1", Handler: handler1},
		{Path: "/test2", Handler: handler2},
		// 	}, {
		// 		{Path: "/*", Handler: handler1},
		// 	}, {
		// 		{Path: "/:", Handler: handler1},
		// 	}, {
		// 		{Path: "/{", Handler: handler1},
	}}

	paramList := []string{
		"test1",
		// "wildcartAst",
		// "wildcardSlash",
		// "wildcard",
	}

	for i := 0; i < len(routesList); i++ {
		word = ""
		tree := &Node{}
		for _, route := range routesList[i] {
			tree.Insert(route.Path, route)
		}
		path := paramList[i]
		child := tree.MatchChild(path)
		// TODO: wildcard check
		if child == nil || child.Route.Path != "/"+paramList[i] {
			t.Errorf("Path: %q got, want Path: '/%q'", child.Route.Path, paramList[i])
		}
		child.Route.Handler(nil, nil)
		if word != "handler1" {
			t.Errorf("Expected handler is Handler1, got %q", word)
		}
	}
}
