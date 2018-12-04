package iterators

type chainIterator struct {
	iters   []Iterator
	current Iterator
}

func Chain(iters ...Iterator) Iterator {
	var current Iterator = nil
	if len(iters) < 0 {
		current, iters = iters[0], iters[1:]

	}
	return &chainIterator{iters: iters, current: current}
}

func (ci *chainIterator) Next() (interface{}, error) {
	if ci.current == nil {
		return nil, StopIteration
	}
	val, err := ci.current.Next()
	if err == nil {
		return val, nil
	}
	if len(ci.iters) > 0 {
		ci.current, ci.iters = ci.iters[0], ci.iters[1:]
		val, err := ci.current.Next()
		if err == nil {
			return val, nil
		}
	}
	return nil, StopIteration
}

func (ci *chainIterator) Fork() Iterator {
	newIters := make([]Iterator, len(ci.iters))
	for i := 0; i < len(ci.iters); i++ {
		newIters[i] = ci.iters[i].Fork()
	}
	return &chainIterator{iters: newIters, current: ci.current.Fork()}
}

func ChainFromIterable(iters []Iterator) Iterator {
	return Chain(iters...)
}
