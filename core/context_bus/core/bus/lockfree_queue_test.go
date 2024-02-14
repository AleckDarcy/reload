package bus

import (
	"runtime"
	"sync"
	"testing"
)

func TestLockFreeQueue(t *testing.T) {
	q := NewLockFreeQueue()
	m := map[int]struct{}{}
	l := sync.Mutex{}
	n := 100

	wg := sync.WaitGroup{}

	eq := func(q *LockFreeQueue, start, num int) {
		for i := 0; i < num; i++ {
			q.Enqueue(start + i)
			runtime.Gosched()
		}

		wg.Done()
	}

	dq := func(q *LockFreeQueue, num int, wait, print bool) {
		for i := 0; i < num; i++ {
			v, ok := q.Dequeue()
			if !ok {
				t.Error()
			}

			if v != nil {
				l.Lock()
				m[v.(int)] = struct{}{}
				l.Unlock()
				if print {
					t.Log(v)
				}
			}
		}

		if wait {
			wg.Done()
		}
	}

	// case 1: n threads write, n threads read
	for i := 0; i < n; i++ {
		wg.Add(1)
		go eq(q, i*n, n)
	}
	wg.Wait()

	if q.Length() != n*n {
		t.Error()
	}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go dq(q, n, true, false)
	}
	wg.Wait()

	if len(m) != n*n {
		t.Error()
	}

	for i := 0; i < n*n; i++ {
		if _, ok := m[i]; !ok {
			t.Error()
		}
	}

	// case 1: n threads write, 1 thread reads
	m = map[int]struct{}{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go eq(q, i*n, n)
	}
	wg.Wait()

	dq(q, n*n, false, true)
	if len(m) != n*n {
		t.Error()
	}

	for i := 0; i < n*n; i++ {
		if _, ok := m[i]; !ok {
			t.Error()
		}
	}
}
