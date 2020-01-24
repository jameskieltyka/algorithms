package recordorder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_orders_recordOrder(t *testing.T) {

	tests := []struct {
		name        string
		o           *orders
		val         int
		expectedPos int
	}{
		{"add order to start", &orders{
			nextPos: 0,
			buffer:  make([]int, 2, 2),
		}, 1, 1},
		{"add order to start", &orders{
			nextPos: 1,
			buffer:  make([]int, 2, 2),
		}, -1, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.o.recordOrder(tt.val)
			assert.Equal(t, tt.o.nextPos, tt.expectedPos)
			assert.Equal(t, tt.val, tt.o.buffer[(tt.expectedPos-1+len(tt.o.buffer))%(len(tt.o.buffer))])
		})
	}
}

func Test_orders_getLast(t *testing.T) {

	tests := []struct {
		name string
		o    *orders
		pos  int
		want int
	}{
		{"get prev index without wrapping", &orders{
			nextPos: 1,
			buffer:  []int{1, 2},
		}, 0, 1},
		{"get prev index with wrapping", &orders{
			nextPos: 1,
			buffer:  []int{1, 2},
		}, 1, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.o.getLast(tt.pos)
			assert.Equal(t, tt.want, got)
		})
	}
}
