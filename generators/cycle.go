package generators

import "github.com/danielSanchezQ/go-iter/iterators"

type cycleIterator struct {
	fork iterators.Iterator
	iter iterators.Iterator
}

func Cycle(iter iterators.Iterator) iterators.Iterator {
	return &cycleIterator{fork: iter.Fork(), iter: iter}
}

func (ci *cycleIterator) Next() (interface{}, error) {
	val, err := ci.iter.Next()
	if err != nil {
		ci.iter = ci.fork.Fork()
		if val, err = ci.iter.Next(); err != nil {
			return nil, err
		}
	}
	return val, nil
}

func (ci *cycleIterator) Fork() iterators.Iterator {
	return &cycleIterator{fork: ci.fork.Fork(), iter: ci.iter.Fork()}
}
