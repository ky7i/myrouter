package myrouter

import (
	"net/http"
	"testing"
)

func TestAdd(t *testing.T) {
	word := ""
	handler1 := func(_ http.ResponseWriter, _ *http.Request) {
		word = "handler1"
	}
	// handler2 := func(_ http.ResponseWriter, _ *http.Request) {
	// 	word = "handler2"
	// }

	tree := &Tree{}

	tree.add("/test", handler1)

	node := tree.get("/test")
	// TODO: how to check path
	// if node.path != "/test" {
	// 	t.Errorf("Add path was %q, want '/test'", node.path)
	// }
	node.Handler(nil, nil)
	if word != "handler1" {
		t.Errorf("Add handler is different one.")
	}
	if node.Path != "/test" {
		t.Errorf("Add Path is %q, want '/test'", node.Path)
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
// 	tree.
// }
