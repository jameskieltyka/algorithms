package recordorder

type orders struct {
	nextPos int
	buffer  []int
}

func (o *orders) recordOrder(id int) {
	o.buffer[o.nextPos] = id
	o.nextPos = (o.nextPos + 1) % len(o.buffer)
}

func (o *orders) getLast(i int) int {
	pos := (o.nextPos - 1 - i + len(o.buffer)) % len(o.buffer)
	return o.buffer[pos]
}
