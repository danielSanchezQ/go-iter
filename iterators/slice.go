package iterators

type sliceIterator struct {
	currentIdx int
	sliceRef   []interface{}
}

func SliceIter(sliceRef []interface{}) Iterator {
	return &sliceIterator{
		sliceRef: sliceRef,
	}
}

func (iter *sliceIterator) Next() (res interface{}, err error) {
	if iter.currentIdx > len(iter.sliceRef) {
		return nil, StopIteration
	}
	res = iter.sliceRef[iter.currentIdx]
	iter.currentIdx++
	return res, nil
}

func (iter *sliceIterator) Fork() Iterator {
	return &sliceIterator{currentIdx: iter.currentIdx, sliceRef: iter.sliceRef}
}
