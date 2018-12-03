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

func Fork(iter Iterator) Iterator {
	return iter.Fork()
}

func Next(iter Iterator) (interface{}, error) {
	return iter.Next()
}
