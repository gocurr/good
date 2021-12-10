package streaming

import (
	"github.com/gocurr/partition"
	"runtime"
	"sort"
	"sync"
)

var cpu = runtime.NumCPU()

var parallelEmpty = &ParallelStream{}

type ParallelStream struct {
	*Stream
	parts []Slice
	wg    sync.WaitGroup
	mu    sync.Mutex
}

// ParallelOf wraps input into *ParallelStream
//
// Returns parallelEmpty when raw is nil
// Or is NOT a slice or an array
func ParallelOf(raw interface{}) *ParallelStream {
	stream := Of(raw)
	slice := stream.slice

	return &ParallelStream{
		Stream: stream,
		parts:  split(slice),
		wg:     sync.WaitGroup{},
		mu:     sync.Mutex{},
	}
}

// split Slice into two-dimensional
func split(slice Slice) []Slice {
	var parts []Slice
	ranges := partition.RangesN(len(slice), cpu)
	for _, r := range ranges {
		parts = append(parts, slice[r.From:r.To])
	}
	return parts
}

// ForEach performs an action for each element of this stream
// in a Parallel way
func (s *ParallelStream) ForEach(act func(interface{})) {
	for _, part := range s.parts {
		s.wg.Add(1)
		slice := part
		go func() {
			for _, v := range slice {
				act(v)
			}
			s.wg.Done()
		}()
	}
	s.wg.Wait()
}

// ForEachOrdered performs an action in order for each element of this stream.
func (s *ParallelStream) ForEachOrdered(act func(interface{})) {
	s.Stream.ForEach(act)
}

// Map returns a stream consisting of the results of applying the given
// function to the elements of this stream in a Parallel way
func (s *ParallelStream) Map(apply func(interface{}) interface{}) *ParallelStream {
	if len(s.slice) == 0 {
		return parallelEmpty
	}

	var mapSlice = make(map[int]Slice)
	for i, part := range s.parts {
		s.wg.Add(1)
		part := part
		go func(i int) {
			var slice Slice
			for _, v := range part {
				slice = append(slice, apply(v))
			}
			s.mu.Lock()
			mapSlice[i] = slice
			s.mu.Unlock()
			s.wg.Done()
		}(i)
	}
	s.wg.Wait()

	var slice Slice
	for i := 0; i < cpu; i++ {
		for _, e := range mapSlice[i] {
			slice = append(slice, e)
		}
	}

	return &ParallelStream{
		Stream: &Stream{slice: slice},
		parts:  split(slice),
	}
}

// Reduce performs a reduction on the elements of this stream,
// using the provided comparing function in a Parallel way
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

	t := vs[0]
	for i := 1; i < len(vs); i++ {
		ti := vs[i]
		if compare(ti, t) {
			t = ti
		}
	}

	return t
}

// Sum returns the sum of elements in this stream
// using the provided sum function in a Parallel way
func (s *ParallelStream) Sum(sum func(interface{}) float64) float64 {
	var rr float64
	for _, _slice := range s.parts {
		s.wg.Add(1)
		slice_ := _slice
		go func() {
			var r float64
			for _, v := range slice_ {
				r += sum(v)
			}
			s.mu.Lock()
			rr += r
			s.mu.Unlock()
			s.wg.Done()
		}()
	}
	s.wg.Wait()
	return rr
}

// Copy returns a new stream containing the elements,
// the new stream holds a copied slice
func (s *ParallelStream) Copy() *ParallelStream {
	slice := make(Slice, len(s.slice))
	copy(slice, s.slice)
	return &ParallelStream{
		Stream: &Stream{slice: slice},
		parts:  split(slice),
	}
}

// Sorted returns a sorted stream consisting of the elements of this stream
// sorted according to the provided less.
//
// Sorted reorders inside slice
// For keeping the order relation of original slice, use Copy first
func (s *ParallelStream) Sorted(less func(i, j int) bool) *ParallelStream {
	sort.Slice(s.slice, less)
	s.parts = split(s.slice)
	return s
}
