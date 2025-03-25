package service

import "sync/atomic"

func atomicCounter(start int64) func() int64 {
	var counter int64 = start
	return func() int64 {
		return atomic.AddInt64(&counter, 1)
	}
}

var Next func() int64

func MakeCounter(start int64) {
	Next = atomicCounter(start)
}
