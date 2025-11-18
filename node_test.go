package myrouter

// func TestMatchChildMatched(t *testing.T) {
// 	word := ""
// 	handler1 := func(_ http.ResponseWriter, _ *http.Request) { word = "handler1" }
// 	handler2 := func(_ http.ResponseWriter, _ *http.Request) { word = "handler2" }
//
// 	routesList := [][]Route{{
// 		{Path: "/test1", Handler: handler1},
// 		{Path: "/test2", Handler: handler2},
// 		// 	}, {
// 		// 		{Path: "/*", Handler: handler1},
// 		// 	}, {
// 		// 		{Path: "/:", Handler: handler1},
// 		// 	}, {
// 		// 		{Path: "/{", Handler: handler1},
// 	}}
//
// 	paramList := []string{
// 		"test1",
// 		// "wildcartAst",
// 		// "wildcardSlash",
// 		// "wildcard",
// 	}
//
// 	for i := 0; i < len(routesList); i++ {
// 		word = ""
// 		tree := &Node{}
// 		for _, route := range routesList[i] {
// 			tree.Insert(route.Path, route)
// 		}
// 		part := paramList[i]
// 		child := tree.MatchChild(part)
// 		// TODO: wildcard check
// 		if child == nil || child.Route.Path != "/"+part {
// 			t.Errorf("Path: %q got, want Path: '/%q'", child.Route.Path, part)
// 		}
// 		child.Route.Handler(nil, nil)
// 		if word != "handler1" {
// 			t.Errorf("Expected handler is Handler1, got %q", word)
// 		}
// 	}
// }

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
