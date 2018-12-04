package iterators

type dropIterator struct {
	iter Iterator
	n    int
}

func Drop(n int, iter Iterator) Iterator {
	return &dropIterator{iter: iter, n: n}
}

func (ti *dropIterator) Next() (interface{}, error) {
	for i := 0; i < ti.n; i++ {
		if _, err := ti.iter.Next(); err != nil {
			return nil, StopIteration
		}
	}
	val, err := ti.iter.Next()
	if err != nil {
		return nil, StopIteration
	}
	return val, nil
}

func (ti *dropIterator) Fork() Iterator {
	return &dropIterator{iter: ti.iter.Fork(), n: ti.n}
}

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
