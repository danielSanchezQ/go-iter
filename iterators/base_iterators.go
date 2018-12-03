package iterators

import "github.com/pkg/errors"

var (
	StopIteration = errors.New("Iterator exhausted")
)

type Predicate func(interface{}) bool

type Iterator interface {
	Next() (interface{}, error)
	Fork() Iterator
}

type sliceIterator struct {
	currentIdx int
	sliceRef   []interface{}
}

func Fork(iter Iterator) Iterator {
	return iter.Fork()
}

func Next(iter Iterator) (interface{}, error) {
	return iter.Next()
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

type zipIterator struct {
	iter1 Iterator
	iter2 Iterator
}

func Zip(iter1, iter2 Iterator) Iterator {
	return &zipIterator{iter1, iter2}
}

func (zip *zipIterator) Next() (interface{}, error) {
	val1, err1 := zip.iter1.Next()
	val2, err2 := zip.iter2.Next()
	if err1 != nil || err2 != nil {
		return nil, StopIteration
	}
	val := [2]interface{}{val1, val2}
	return val, nil
}

func (zip *zipIterator) Fork() Iterator {
	iter1 := zip.iter1.Fork()
	iter2 := zip.iter2.Fork()
	return Zip(iter1, iter2)
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
