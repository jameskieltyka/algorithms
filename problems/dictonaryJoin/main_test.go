package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordSets(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		dict []string
		word string
		want [][]string
	}{
		{"test a", []string{"the", "quick", "brown", "fox"}, "thequickbrownfox", [][]string{{"the", "quick", "brown", "fox"}}},
		{"test b", []string{"bed", "bath", "bedbath", "and", "beyond"}, "bedbathandbeyond",
			[][]string{{"bed", "bath", "and", "beyond"}, {"bedbath", "and", "beyond"}}},
		{"test c", []string{"bed", "bath", "bedbath", "and", "beyond"}, "bedbathaandbeyond", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordSets(tt.dict, tt.word)
			assert.Equal(t, tt.want, got)
		})
	}
}
