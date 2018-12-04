package iterators

type takeIterator struct {
	iter    Iterator
	current int
	n       int
}

func Take(n int, iter Iterator) Iterator {
	return &takeIterator{iter: iter, current: 0, n: n}
}

func (ti *takeIterator) Next() (interface{}, error) {
	if ti.current > ti.n {
		return nil, StopIteration
	}
	val, err := ti.iter.Next()
	if err != nil {
		return nil, StopIteration
	}
	ti.current++
	return val, nil
}

func (ti *takeIterator) Fork() Iterator {
	return &takeIterator{iter: ti.iter.Fork(), current: ti.current, n: ti.n}
}

type takeWhileIterator struct {
	f           Predicate
	iter        Iterator
	stopReached error
}

func TakeWhile(f Predicate, iter Iterator) Iterator {
	return &takeWhileIterator{f, iter, nil}
}

func (tw *takeWhileIterator) Next() (val interface{}, err error) {
	val, err = tw.iter.Next()
	if err != nil || !tw.f(val) || tw.stopReached != nil {
		tw.stopReached = StopIteration
		return nil, StopIteration
	}
	return val, nil
}

func (tw *takeWhileIterator) Fork() Iterator {
	return TakeWhile(tw.f, tw.iter.Fork())
}
