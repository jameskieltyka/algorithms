package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubarray(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name     string
		inputArr []int
		window   int
		want     []int
	}{
		{"test a", []int{10, 5, 2, 7, 8, 7}, 3, []int{10, 7, 8, 8}},
		{"test b", []int{10, 5, 6, 3, 9, 4, 1}, 4, []int{10, 9, 9, 9}},
		{"test b", []int{1, -1, -2, -3}, 2, []int{1, -1, -2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubarray(tt.inputArr, tt.window)
			assert.Equal(t, tt.want, got)
		})
	}
}
