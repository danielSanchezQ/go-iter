package generators

import "github.com/danielSanchezQ/go-iter/iterators"

type countIterator struct {
	i    int
	step int
}

func Count(start, step int) iterators.Iterator {
	return &countIterator{i: start, step: step}
}

func (ci *countIterator) Next() (interface{}, error) {
	val := ci.i
	ci.i += ci.step
	return val, nil
}

func (ci *countIterator) Fork() iterators.Iterator {
	return Count(ci.i, ci.step)
}
