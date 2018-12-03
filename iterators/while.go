package iterators

type dropWhileIterator struct {
	f    Predicate
	iter Iterator
}

func DropWhile(f Predicate, iter Iterator) Iterator {
	return &dropWhileIterator{f, iter}
}

func (dw *dropWhileIterator) Next() (val interface{}, err error) {
	for val, err = dw.iter.Next(); dw.f(val) && err != nil; {

	}
	if err != nil {
		return nil, StopIteration
	}
	return val, nil
}

func (dw *dropWhileIterator) Fork() Iterator {
	return DropWhile(dw.f, dw.iter.Fork())
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
