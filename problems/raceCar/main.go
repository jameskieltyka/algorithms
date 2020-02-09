package main

import "fmt"

//Leetcode 818
func main() {
	fmt.Println(racecar(99899))
}

func racecar(target int) int {

	queue := make(chan loc, 10000000)
	posMap := make(map[string]bool)
	queue <- loc{
		speed: 1,
		pos:   0,
		count: 0,
	}

	for len(queue) > 0 {
		l := <-queue
		if l.pos == target {
			return l.count
		}

		if posMap[l.string()] {
			continue
		}

		posMap[l.string()] = true

		a := l
		a.accelerate()
		if a.pos < (target<<1) && a.pos > 0 {
			queue <- a
		}

		r := l
		r.reverse()
		if r.pos < (target<<1) && r.pos > 0 {
			queue <- r
		}
	}

	return 0
}

type loc struct {
	speed int
	pos   int
	count int
}

func (l *loc) reverse() {
	l.count++
	if l.speed > 0 {
		l.speed = -1
		return
	}
	l.speed = 1
	return
}

func (l *loc) accelerate() {
	l.pos = l.pos + l.speed
	l.speed = l.speed * 2
	l.count++
}

func (l *loc) string() string {
	return fmt.Sprintf("%s,%s", l.pos, l.speed)
}
