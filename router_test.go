package myrouter

import (
	"net/http"
	"testing"
)

func dummyHandler(_ http.ResponseWriter, _ *http.Request) {}

func TestNew(t *testing.T) {
	r := New()
	length := len(r.trees)
	if length != 10 {
		t.Errorf("Router.trees had %q elements, want %q.", length, 10)
	}
}

func TestHTTPMethods(t *testing.T) {
	r := New()
	r.GET("/get", dummyHandler)
	r.POST("/post", dummyHandler)

	methodIndexs := []int{
		getMethodIndexOf("GET"),
		getMethodIndexOf("POST"),
	}

	for _, index := range methodIndexs {
		if len(r.trees[index].root.Children) != 1 {
			t.Errorf("A given path didn't be registered at a correct Method tree.")
		}
	}
}

func TestInsert(t *testing.T) {
	dummyHandler := func(_ http.ResponseWriter, _ *http.Request) {}

	r := New()

	invalidPaths := []string{"", "noSlash"}
	for _, path := range invalidPaths {
		if err := catchPanic(func() { r.insert("GET", path, dummyHandler) }); err == nil {
			t.Errorf("Didn't throw a panic to this invalid path %q.", path)
		}
	}

	r = New()
	paths := []string{
		"/api",
		"/users",
	}
	for _, path := range paths {
		r.insert("GET", path, dummyHandler)
	}

	insertCount := len(paths)
	methodGETLen := len(r.trees[0].root.Children) // trees[0] is the GET method tree

	if methodGETLen != insertCount {
		t.Errorf("%d GET methods were registerd, want %d.", methodGETLen, insertCount)
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
