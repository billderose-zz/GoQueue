package queue

type Queue chan int

func (q Queue) Enqueue(i int) {
	q <- i
}

func (q Queue) Dequeue() (int, bool) {
	val, err := <-q
	return val, err
}

func (q Queue) Close() {
	close(q)
}
