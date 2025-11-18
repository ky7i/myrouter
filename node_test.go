package myrouter

import (
	"testing"
)

func TestMatchChildMatched(t *testing.T) {
	// 	word := ""
	// 	handler1 := func(_ http.ResponseWriter, _ *http.Request) { word = "handler1" }
	// 	handler2 := func(_ http.ResponseWriter, _ *http.Request) { word = "handler2" }

	nodes := []*Node{
		{Part: "a"},
		{Part: "aa"},
		{Part: "A"},
		{Part: "b"},
	}

	node := &Node{
		Children: nodes,
	}

	paramList := []string{
		"a",
		// "wildcartAst",
		// "wildcardSlash",
		// "wildcard",
	}

	for _, param := range paramList {
		// word = ""
		child := node.MatchChild(param)
		// TODO: wildcard check
		if child == nil || child.Part != param {
			t.Errorf("Got part is nil, want %q", param)
		} else if child.Part != param {
			t.Errorf("Got part is %q, want %q", child.Part, param)
		}
		//		child.Route.Handler(nil, nil)
		//		if word != "handler1" {
		//			t.Errorf("Expected handler is Handler1, got %q", word)
		//		}
	}
}

// func TestSearch(t *testing.T) {
// 	routes := [][]&Routes{
// 		{
// 			{Path: "/test1/test2", Handler: handler1},
// 		}, {
// 			{Path:}
// 		}
// 	}
//
// 	paramList := []string{
// 		"/test1/test2",
//
// 	}
//
// 	for i := 0; i < len(routesList); i++ {
//
// 	}
//
// }
