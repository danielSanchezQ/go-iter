package iterators

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
