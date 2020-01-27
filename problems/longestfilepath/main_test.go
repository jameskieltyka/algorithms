package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPath(t *testing.T) {
	type args struct {
		fs string
	}
	tests := []struct {
		name string
		fs   string
		want int
	}{
		{"test 1", "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", 20},
		{"test 2", "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
		{"test 3", "dir\n\tsubdir1", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPath(tt.fs)
			assert.Equal(t, tt.want, got)
		})
	}
}
