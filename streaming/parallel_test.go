package streaming

import (
	"fmt"
	"testing"
)

var p = ParallelOf([]int{1, 1, 3, 3, 6, 6, 7})

func TestStream_ForEach(t *testing.T) {
	p.ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func TestParallelStream_Map(t *testing.T) {
	slice := p.Map(func(i interface{}) interface{} {
		return i.(int) * 2
	}).Collect()
	fmt.Printf("%v\n", slice)
}

func TestParallelStream_Reduce(t *testing.T) {
	reduce := p.Reduce(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})
	fmt.Printf("%v\n", reduce)
}

func TestParallelStream_Distinct(t *testing.T) {
	slice := p.Distinct().Collect()
	fmt.Printf("%v\n", slice)
}

func TestParallelStream_Sum(t *testing.T) {
	sum := p.Sum(func(i interface{}) float64 {
		return float64(i.(int))
	})
	fmt.Printf("%v\n", sum)
}
