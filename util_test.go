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
	paths := []string{
		"/users",
		"/a",
		"/",
		"//",
	}

	for _, path := range paths {
		catchPanic(func() { validatePath(path) })
	}

}
