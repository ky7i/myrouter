package myrouter

import (
	"net/http"
	"slices"
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

	// check if retrun nil when mismatch
	if node := tree.get("/mismatchPath"); node != nil {
		t.Errorf("want nil, got %q.", node.Part)
	}

	invalidPaths := []string{
		"/apis//multipleSlash",
		"/combination//",
	}

	for _, path := range invalidPaths {
		if err := catchPanic(func() { tree.add(path, fakeHandler(path)) }); err == nil {
			t.Errorf("Panic didn't be thrown.")
		}
	}
}

func TestSplitPath(t *testing.T) {
	paths := []string{
		"/test",
		"/user/userId",
		"//multipleSlash",
		"/trailingSlash/",
		"/",
	}

	partsList := [][]string{
		{"test"},
		{"user", "userId"},
		{"", "multipleSlash"},
		{"trailingSlash"},
		{},
	}

	for i := range paths {
		path := splitPath(paths[i])
		if !slices.Equal(path, partsList[i]) {
			t.Errorf("A splited path was %q, want %q.", path, partsList[i])
		}
	}
}

// test if registerd nodes are rollbacked in panic.
// => not in Golang
// when panic, main gorutine finishes and rollback is unnecessary.
// func TestPanicRollbackInAdd(t *testing.T) {
// 	tree := &Tree{}
// 	catchPanic(func() { tree.add("/invalidPath//", func(_ http.ResponseWriter, _ *http.Request) {}) })
//
// 	node := tree.get("/invalidPath")
// 	if node != nil {
// 		t.Errorf("Did not be rollbacked.")
// 	}
// }
