package myrouter

import (
	"net/http"
	"strings"
	"testing"
)

var fakeHandlerValue string

func fakeHandler(str string) (h http.HandlerFunc) {
	return func(http.ResponseWriter, *http.Request) {
		fakeHandlerValue = str
	}
}

func TestAdd(t *testing.T) {
	tree := &Tree{root: &Node{}}

	correctPaths := []string{
		"/test",
		"/apis/users/userId",
	}

	for _, path := range correctPaths {
		tree.add(path, fakeHandler(path))
	}

	for _, path := range correctPaths {
		node := tree.get(path)
		lastChild := path[(strings.LastIndex(path, "/") + 1):]

		if node.Part != lastChild {
			t.Errorf("A deepest node was %q, want to be %q", node.Part, lastChild)
		}

		node.Handler(nil, nil)
		if fakeHandlerValue != path {
			t.Errorf("A set Handler was wrong, got %q, want %q", fakeHandlerValue, path)
		}
	}

	invalidPaths := []string{
		"/trailingSlash/",
		"/apis//multipleSlash",
		"/combination//",
	}

	for _, path := range invalidPaths {
		if err := catchPanic(func() { tree.add(path, fakeHandler(path)) }); err == nil {
			t.Errorf("Panic didn't be thrown.")
		}
	}
}

// func TestGet(t *testing.T) {
// 	word := ""
// 	handler1 := func(_ http.ResponseWriter, _ *http.Request) {
// 		word = "handler1"
// 	}
//
// 	tree := &Tree{}
//
// }
