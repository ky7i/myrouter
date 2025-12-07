package myrouter

import (
	"testing"
)

func catchPanic(testFunc func()) (recv interface{}) {
	defer func() {
		recv = recover()
	}()
	testFunc()
	return
}

func TestValidatePath(t *testing.T) {
	correctPaths := []string{
		"/users",
		"/a",
	}

	for _, path := range correctPaths {
		if err := catchPanic(func() { validatePath(path) }); err != nil {
			t.Errorf("Got an enexpected error, %q", err)
		}
	}

	incorrectPaths := []string{
		"a",
	}

	for _, path := range incorrectPaths {
		if err := catchPanic(func() { validatePath(path) }); err == nil {
			t.Errorf("This method didn't call a panic.")
		}
	}
}
