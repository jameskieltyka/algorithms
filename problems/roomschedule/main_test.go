package main

import "testing"

func Test_minimumRooms(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name     string
		bookings []booking
		want     int
	}{
		{"test a", []booking{{30, 70}, {0, 50}, {60, 150}}, 2},
		{"test b", []booking{{30, 70}, {0, 50}, {60, 150}, {50, 70}, {90, 160}}, 3},
		{"test c", []booking{{10, 70}, {0, 50}, {30, 150}, {50, 70}, {10, 160}}, 4},
		{"test d", []booking{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumRooms(tt.bookings); got != tt.want {
				t.Errorf("minimumRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
