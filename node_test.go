package myrouter

import (
	"testing"
)

func TestMatchChildMatched(t *testing.T) {

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
	}

	for _, param := range paramList {
		child := node.MatchChild(param)
		if child == nil {
			t.Errorf("Got part is nil, want %q", param)
		} else if child.Part != param {
			t.Errorf("Got part is %q, want %q", child.Part, param)
		}
	}
}
