package myrouter

import (
	"net/http"
	"testing"
)

func TestNew(t *testing.T) {
	r := New()
	length := len(r.trees)
	if length != 10 {
		t.Errorf("Router.trees had %q elements, want %q.", length, 10)
	}
}

func TestInsert(t *testing.T) {
	fakeHandler := func(_ http.ResponseWriter, _ *http.Request) {}

	r := New()
	r.insert("GET", "/api/users", fakeHandler)
	r.insert("GET", "/api/sample", fakeHandler)

	methodGETLen := len(r.trees[0].root.Children)

	if methodGETLen != 2 {
		t.Errorf("%q GET methods were registerd, want %q.", methodGETLen, 2)
	}
}
