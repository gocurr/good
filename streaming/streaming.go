package streaming

import (
	"reflect"
)

// empty stream
var empty = &Stream{}

// Slice alias of interface slice
type Slice []interface{}

// Stream Slice holder
type Stream struct {
	slice Slice
}

// Of wraps input into *Stream
// Note: if input is nil, returns empty
// Note: if input is not a slice or an array, returns empty
func Of(raw interface{}) *Stream {
	if raw == nil {
		return empty
	}

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

// FlatMap returns a stream consisting of the results
// of replacing each element of this stream
func (s *Stream) FlatMap(apply func(interface{}) interface{}) *Stream {
	stream := s.Map(apply)
	slice := stream.slice
	if len(slice) == 0 {
		return empty
	}

	switch reflect.TypeOf(slice[0]).Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return stream
	}

	var r Slice
	for _, _slice := range slice {
		value := reflect.ValueOf(_slice)
		for i := 0; i < value.Len(); i++ {
			ele := value.Index(i)
			r = append(r, ele.Interface())
		}
	}
	return &Stream{slice: r}
}

// Peek returns the same stream,
// additionally performing the provided action on each element
// as elements are consumed from the resulting stream
func (s *Stream) Peek(act func(interface{})) *Stream {
	for _, v := range s.slice {
		act(v)
	}
	return s
}

// Limit returns a stream consisting of the elements of this stream,
// truncated to be no longer than max-size in length.
func (s *Stream) Limit(n int) *Stream {
	if n <= 0 {
		return empty
	}
	if len(s.slice) == 0 {
		return empty
	}

	if n >= len(s.slice) {
		return s
	}

	return &Stream{slice: s.slice[:n]}
}

// Skip returns a stream consisting of the remaining elements
// of this stream after discarding the first n elements
// of the stream. If the stream contains fewer than n elements then
// an empty stream will be returned.
func (s *Stream) Skip(n int) *Stream {
	if n <= 0 {
		return s
	}
	if len(s.slice) == 0 {
		return empty
	}

	if n >= len(s.slice) {
		return empty
	}

	return &Stream{slice: s.slice[n:]}
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

	for j := 1; j < len(s.slice); j++ {
		v := s.slice[j]
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

// FilterCount returns count of the elements of this stream
// that match the given predicate.
func (s *Stream) FilterCount(predicate func(interface{}) bool) int {
	var c int
	for _, v := range s.slice {
		if predicate(v) {
			c++
		}
	}
	return c
}

// Distinct returns a stream consisting of the distinct elements
// with original order
func (s *Stream) Distinct() *Stream {
	if len(s.slice) == 0 {
		return empty
	}

	var slice []interface{}

	var memory = make(map[interface{}]int)
	for i, v := range s.slice {
		if _, ok := memory[v]; !ok {
			memory[v] = i
			slice = append(slice, v)
		}
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

// IsEmpty reports stream is empty
func (s *Stream) IsEmpty() bool {
	return len(s.slice) == 0
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

// AnyMatch returns whether any elements of this stream match
// the provided predicate. May not evaluate the predicated
// on all elements if not necessary for determining the result.
// If the stream is empty then false is returned and the predicate is not evaluated.
func (s *Stream) AnyMatch(predicate func(interface{}) bool) bool {
	if len(s.slice) == 0 {
		return false
	}

	for _, v := range s.slice {
		if predicate(v) {
			return true
		}
	}
	return false
}

// AllMatch returns whether all elements of this stream match
// the provided predicate. May not evaluate the predicated
// on all elements if not necessary for determining the result.
// If the stream is empty then true is returned and the predicate is not evaluated.
func (s *Stream) AllMatch(predicate func(interface{}) bool) bool {
	if len(s.slice) == 0 {
		return true
	}

	for _, v := range s.slice {
		if predicate(v) {
			continue
		}
		return false
	}
	return true
}

// NoneMatch returns whether no elements of this stream match
// the provided predicate. May not evaluate the predicated
// on all elements if not necessary for determining the result.
// If the stream is empty then true is returned and the predicate is not evaluated.
func (s *Stream) NoneMatch(predicate func(interface{}) bool) bool {
	if len(s.slice) == 0 {
		return true
	}

	for _, v := range s.slice {
		if !predicate(v) {
			continue
		}
		return false
	}
	return true
}

// FindFirst returns the first element of the stream,
// or nil if the stream is empty
func (s *Stream) FindFirst() interface{} {
	if len(s.slice) == 0 {
		return nil
	}
	return s.slice[0]
}
