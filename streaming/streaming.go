package streaming

import (
	"reflect"
)

var (
	nothing struct{}
	empty   = &Stream{}
)

// Slice alias of interface slice
type Slice []interface{}

// Stream Slice holder
type Stream struct {
	slice Slice
}

// Of wraps input into *Stream
// Note: if input is not a slice or an array, returns empty
func Of(raw interface{}) *Stream {
	switch reflect.TypeOf(raw).Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return empty
	}

	var slice Slice
	value := reflect.ValueOf(raw)
	for i := 0; i < value.Len(); i++ {
		ele := value.Index(i)
		slice = append(slice, ele.Interface())
	}

	return &Stream{slice: slice}
}

// ForEach performs an action for each element of this stream.
func (s *Stream) ForEach(act func(interface{})) {
	for _, v := range s.slice {
		act(v)
	}
}

// Map returns a stream consisting of the results of applying the given
// function to the elements of this stream.
func (s *Stream) Map(apply func(interface{}) interface{}) *Stream {
	if len(s.slice) == 0 {
		return empty
	}

	var slice Slice
	for _, v := range s.slice {
		slice = append(slice, apply(v))
	}
	return &Stream{slice: slice}
}

// Limit returns a stream consisting of the elements of this stream,
// truncated to be no longer than max-size in length.
func (s *Stream) Limit(n int) *Stream {
	if len(s.slice) == 0 {
		return empty
	}

	length := len(s.slice)
	if n > length {
		n = length
	}
	return &Stream{slice: s.slice[:n]}
}

// Reduce performs a reduction on the elements of this stream,
// using the provided comparing function
// NOTE: when steam is empty, Reduce returns -1 as the index
func (s *Stream) Reduce(compare func(a, b interface{}) bool) (interface{}, int) {
	if len(s.slice) == 0 {
		return nil, -1
	}

	t := s.slice[0]
	i := 0
	for j, v := range s.slice {
		if compare(v, t) {
			t = v
			i = j
		}
	}
	return t, i
}

// Filter returns a stream consisting of the elements of this stream
// that match the given predicate.
func (s *Stream) Filter(predicate func(interface{}) bool) *Stream {
	if len(s.slice) == 0 {
		return empty
	}

	var slice Slice
	for _, v := range s.slice {
		if predicate(v) {
			slice = append(slice, v)
		}
	}
	return &Stream{slice: slice}
}

// Distinct returns a stream consisting of the distinct elements
func (s *Stream) Distinct() *Stream {
	if len(s.slice) == 0 {
		return empty
	}

	var m = make(map[interface{}]struct{})
	for _, v := range s.slice {
		m[v] = nothing
	}

	var slice Slice
	for k := range m {
		slice = append(slice, k)
	}
	return &Stream{slice: slice}
}

// Collect returns data load of this stream
func (s *Stream) Collect() Slice {
	return s.slice
}

// Count returns the count of elements in this stream
func (s *Stream) Count() int {
	return len(s.slice)
}

// Sum returns the sum of elements in this stream
// using the provided sum function
func (s *Stream) Sum(sum func(interface{}) float64) float64 {
	var r float64
	for _, v := range s.slice {
		r += sum(v)
	}
	return r
}
