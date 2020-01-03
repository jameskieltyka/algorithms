package priority

type Queue struct {
	data []Data
}

type Data struct {
	Key   int
	Value int
}

func New() *Queue {
	return &Queue{
		data: make([]Data, 0, 10),
	}
}

func (q *Queue) Insert(key int, value int) {
	q.data = append(q.data, Data{
		Key:   key,
		Value: value,
	})
	q.reheapUp(q.Len() - 1)
}

func (q *Queue) Len() int {
	return len(q.data)
}

func (q *Queue) Remove() int {
	value := q.data[0].Value
	q.data[0].Key = -1
	q.reheapDown(0)
	q.data = q.data[:q.Len()-1]
	return value
}

func (q *Queue) reheapUp(index int) {
	parent := index / 2
	for q.data[parent].Key < q.data[index].Key && index >= 1 {
		q.data[parent], q.data[index] = q.data[index], q.data[parent]
		index = parent
		parent = parent / 2
	}
}

func (q *Queue) reheapDown(index int) {
	child := 2 * index
	for child <= len(q.data)-1 {

		if child < q.Len()-1 && q.data[child+1].Key > q.data[child].Key {
			child++
		}

		if q.data[index].Key >= q.data[child].Key {
			break
		}

		q.data[index], q.data[child] = q.data[child], q.data[index]

		index = child
		child = 2 * index
	}
}
