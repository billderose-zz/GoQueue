package queue

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const capacity = 10000

func TestQueue(t *testing.T) {
	q := newQueue(capacity)
	var produced, consumed int
	done := make(chan bool)
	go func() {
		for j := 0; j < capacity*10; j++ {
			q.enqueue(rand.Int())
			produced++
			time.Sleep(time.Microsecond)
		}
		q.close()
		done <- true
	}()

	go func() {
		for {
			if _, err := q.dequeue(); !err {
				break
			} else {
				consumed++
			}
			time.Sleep(time.Microsecond)
		}
		done <- true
	}()

	for i := 0; i < 2; i++ {
		<-done
	}

	if produced != consumed {
		t.Error("Enqueue not linearizable; test failed")
	} else {
		t.Log("Test passed")
		fmt.Println(produced)
		fmt.Println(consumed)
	}
}

func BenchmarkQueue(b *testing.B) {
	b.StopTimer()
	var q = newQueue(capacity)
	b.StartTimer()
	for j := 0; j < b.N; j++ {
		q.enqueue(j)
		q.dequeue()
	}
}
