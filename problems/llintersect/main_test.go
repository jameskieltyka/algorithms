package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intersect(t *testing.T) {
	root1 := &node{val: 3}
	root1.next = &node{val: 7}
	root1.next.next = &node{val: 8}
	root1.next.next.next = &node{val: 10}

	root2 := &node{val: 99}
	root2.next = &node{val: 1}
	root2.next.next = &node{val: 8}
	root2.next.next.next = &node{val: 10}

	tests := []struct {
		name  string
		root1 *node
		root2 *node
		want  int
	}{
		{"test a", root1, root2, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intersect(tt.root1, tt.root2)
			assert.Equal(t, tt.want, got)
		})
	}
}
