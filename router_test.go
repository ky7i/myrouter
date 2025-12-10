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
	r.insert("GET", "/api", fakeHandler)
	r.insert("GET", "/users", fakeHandler)

	methodGETLen := len(r.trees[0].root.Children)

	if methodGETLen != 2 {
		t.Errorf("%d GET methods were registerd, want %d.", methodGETLen, 2)
	}
}

func TestGetMethodIndexOf(t *testing.T) {
	methods := []string{"GET", "POST", "INVALID"}
	// TODO: rename, there are similer variable names.
	methodIndexes := []int{0, 1, -1}

	for i := 0; i < len(methods); i++ {
		methodIndex := getMethodIndexOf(methods[i])
		if methodIndex != methodIndexes[i] {
			t.Errorf("Correct index of %q method is %d, set %d", methods[i], methodIndexes[i], methodIndex)
		}
	}
}

// func TestServeHTTP(t *testing.T) {
//
// }
