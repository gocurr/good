package streaming

import (
	"github.com/gocurr/partition"
	"reflect"
	"runtime"
	"sort"
	"sync"
)

var cpu = runtime.NumCPU()

var parallelEmpty = &ParallelStream{}

type ParallelStream struct {
	slice Slice
	wg    sync.WaitGroup
	mu    sync.Mutex
	parts [][]interface{}
}

// ParallelOf wraps input into *ParallelStream
//
// Returns empty when raw is nil
// Or is NOT a slice or an array
func ParallelOf(raw interface{}) *ParallelStream {
	if raw == nil {
		return parallelEmpty
	}

	switch reflect.TypeOf(raw).Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return parallelEmpty
	}

	var slice Slice
	value := reflect.ValueOf(raw)
	for i := 0; i < value.Len(); i++ {
		ele := value.Index(i)
		slice = append(slice, ele.Interface())
	}

	return &ParallelStream{slice: slice, parts: getParts(slice)}
}

func getParts(slice Slice) [][]interface{} {
	var parts [][]interface{}
	ranges := partition.RangesN(len(slice), cpu)
	for _, r := range ranges {
		parts = append(parts, slice[r.From:r.To])
	}
	return parts
}

// ForEach performs an action for each element of this stream.
func (s *ParallelStream) ForEach(act func(interface{})) {
	for _, _slice := range s.parts {
		s.wg.Add(1)
		slice_ := _slice
		go func() {
			for _, v := range slice_ {
				act(v)
			}
			s.wg.Done()
		}()
	}
	s.wg.Wait()
}

// Map returns a stream consisting of the results of applying the given
// function to the elements of this stream.
func (s *ParallelStream) Map(apply func(interface{}) interface{}) *ParallelStream {
	if len(s.slice) == 0 {
		return parallelEmpty
	}

	var mapSlice = make(map[int]Slice)
	for i, part := range s.parts {
		s.wg.Add(1)
		slice_ := part
		go func(i int) {
			var slice Slice
			for _, v := range slice_ {
				slice = append(slice, apply(v))
			}
			s.mu.Lock()
			mapSlice[i] = slice
			s.mu.Unlock()
			s.wg.Done()
		}(i)
	}
	s.wg.Wait()

	var keys = make([]int, 0, len(mapSlice))
	for key := range mapSlice {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	var rSlice Slice
	for _, k := range keys {
		for _, e := range mapSlice[k] {
			rSlice = append(rSlice, e)
		}
	}

	return &ParallelStream{slice: rSlice, parts: getParts(rSlice)}
}

// Reduce performs a reduction on the elements of this stream,
// using the provided comparing function
//
// When steam is empty, Reduce returns nil, -1
func (s *ParallelStream) Reduce(compare func(a, b interface{}) bool) interface{} {
	if len(s.slice) == 0 {
		return nil
	}

	var vs Slice
	for _, part := range s.parts {
		s.wg.Add(1)
		slice := part
		go func() {
			t := slice[0]
			for j := 1; j < len(slice); j++ {
				v := slice[j]
				if compare(v, t) {
					t = v
				}
			}

			s.mu.Lock()
			vs = append(vs, t)
			s.mu.Unlock()

			s.wg.Done()
		}()
	}
	s.wg.Wait()

	tt := vs[0]
	for i := 1; i < len(vs); i++ {
		ti := vs[i]
		if compare(ti, tt) {
			tt = ti
		}
	}

	return tt
}

// Distinct returns a stream consisting of the distinct elements
// with original order
func (s *ParallelStream) Distinct() *ParallelStream {
	if len(s.slice) == 0 {
		return parallelEmpty
	}

	var mapSlice = make(map[int]Slice)
	for i, _slice := range s.parts {
		s.wg.Add(1)
		slice_ := _slice
		go func(i int) {
			var memory = make(map[interface{}]int)
			var slice []interface{}
			for i, v := range slice_ {
				if _, ok := memory[v]; !ok {
					memory[v] = i
					slice = append(slice, v)
				}
			}
			s.mu.Lock()
			mapSlice[i] = slice
			s.mu.Unlock()
			s.wg.Done()
		}(i)
	}
	s.wg.Wait()

	var keys = make([]int, 0, len(mapSlice))
	for key := range mapSlice {
		keys = append(keys, key)
	}
	sort.Ints(keys)
	var rSlice Slice
	for _, k := range keys {
		for _, e := range mapSlice[k] {
			rSlice = append(rSlice, e)
		}
	}

	var r Slice
	var memory = make(map[interface{}]int)
	for i, v := range rSlice {
		if _, ok := memory[v]; !ok {
			memory[v] = i
			r = append(r, v)
		}
	}

	return &ParallelStream{slice: r, parts: getParts(r)}
}

// Collect returns data load of this stream
func (s *ParallelStream) Collect() Slice {
	return s.slice
}

// Count returns the count of elements in this stream
func (s *ParallelStream) Count() int {
	return len(s.slice)
}

// IsEmpty reports stream is empty
func (s *ParallelStream) IsEmpty() bool {
	return len(s.slice) == 0
}

// Sum returns the sum of elements in this stream
// using the provided sum function
func (s *ParallelStream) Sum(sum func(interface{}) float64) float64 {
	var result float64
	for _, _slice := range s.parts {
		s.wg.Add(1)
		slice_ := _slice
		go func() {
			var r float64
			for _, v := range slice_ {
				r += sum(v)
			}
			s.mu.Lock()
			result += r
			s.mu.Unlock()
			s.wg.Done()
		}()
	}
	s.wg.Wait()
	return result
}
