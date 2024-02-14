package bus

import (
	"runtime"
	"sync/atomic"
	"unsafe"
)

type node struct {
	next unsafe.Pointer
	v    interface{}
}

func load(p *unsafe.Pointer) *node {
	return (*node)(atomic.LoadPointer(p))
}
func compareAndSwap(p *unsafe.Pointer, old, new *node) bool {
	return atomic.CompareAndSwapPointer(p, unsafe.Pointer(old), unsafe.Pointer(new))
}

type LockFreeQueue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
	len  int64
}

func NewLockFreeQueue() *LockFreeQueue {
	head := node{}
	return &LockFreeQueue{
		head: unsafe.Pointer(&head),
		tail: unsafe.Pointer(&head),
		len:  0,
	}
}

func (q *LockFreeQueue) Enqueue(v interface{}) {
	i := &node{v: v}
	for {
		tail := load(&q.tail)
		next := load(&tail.next)
		if load(&q.tail) == tail {
			if next == nil {
				if compareAndSwap(&tail.next, next, i) {
					compareAndSwap(&q.tail, tail, i)
					atomic.AddInt64(&q.len, 1)

					return
				}
			} else {
				compareAndSwap(&q.tail, tail, next)
			}
		}

		runtime.Gosched()
	}
}

func (q *LockFreeQueue) Dequeue() (interface{}, bool) {
	for {
		head := load(&q.head)
		tail := load(&q.tail)
		next := load(&head.next)
		if load(&q.head) == head {
			if head == tail {
				if next == nil {
					return nil, false
				}
				compareAndSwap(&q.tail, tail, next)
			} else {
				v := next.v
				if compareAndSwap(&q.head, head, next) {
					atomic.AddInt64(&q.len, -1)

					return v, true
				}
			}
		}

		runtime.Gosched()
	}
}

func (q *LockFreeQueue) Length() int {
	return int(atomic.LoadInt64(&q.len))
}
