package iterators

type filterIterator struct {
	f    Predicate
	iter Iterator
}

func Filter(f Predicate, iter Iterator) Iterator {
	return &filterIterator{f, iter}
}

func (fi *filterIterator) Next() (interface{}, error) {
	for val, err := fi.iter.Next(); err != nil; {
		if fi.f(val) {
			return val, nil
		}
	}
	return nil, StopIteration
}

func (fi *filterIterator) Fork() Iterator {
	return Filter(fi.f, fi.iter.Fork())
}
