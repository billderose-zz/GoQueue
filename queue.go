package queue

type Queue chan int

func newQueue(c int) Queue {
	return make(Queue, c)
}

func (q Queue) enqueue(i int) {
	q <- i
}

func (q Queue) dequeue() (int, bool) {
	val, err := <-q
	return val, err
}

func (q Queue) close() {
	close(q)
}
