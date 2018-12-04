package generators

import "github.com/danielSanchezQ/go-iter/iterators"

type repeatIterator struct {
	value   interface{}
	times   int
	current int
}

func Repeat(val interface{}, times int) iterators.Iterator {
	if times < 0 {
		times = -1
	}
	return &repeatIterator{val, times, 0}
}

func (ri *repeatIterator) Next() (interface{}, error) {
	if ri.times < 0 {
		return ri.value, nil
	}
	if ri.current < ri.times {
		ri.current++
		return ri.value, nil
	}
	return nil, iterators.StopIteration
}

func (ri *repeatIterator) Fork() iterators.Iterator {
	return &repeatIterator{ri.value, ri.times, ri.current}
}
