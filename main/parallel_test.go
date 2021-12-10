package main

import (
	"fmt"
	"github.com/gocurr/good/streaming"
	ss "strings"
	"testing"
)

var p = streaming.ParallelOf(Ints{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 35, 36, 37, 38, 39})

func TestStream_ForEach(t *testing.T) {
	p.ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func TestStream_ForEachOrdered(t *testing.T) {
	p.ForEachOrdered(func(i interface{}) {
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

func TestStream_sum(t *testing.T) {
	sum := p.Stream.Sum(func(i interface{}) float64 {
		return float64(i.(int))
	})
	fmt.Printf("%v\n", sum)
}

func TestParallelCopy(t *testing.T) {
	parallelStream := p.Copy()
	slice := p.Collect()
	p.Sorted(func(i, j int) bool {
		return slice[i].(int) > slice[j].(int)
	})
	mul := p.Map(func(i interface{}) interface{} {
		return i.(int) * 10
	})
	fmt.Printf("%v\n", mul.Collect())
	fmt.Printf("%v\n", parallelStream.Collect())
}

func TestFlatMap(t *testing.T) {
	p := streaming.ParallelOf(strings{"hello there", "good morning", "one", "two", "three", "four", "five"})
	flatMap := p.FlatMap(func(i interface{}) interface{} {
		return ss.Split(i.(string), " ")
	})
	slice := flatMap.Collect()
	fmt.Printf("%v\n", slice)
}