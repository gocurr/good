package main

import (
	"fmt"
	"github.com/gocurr/good/streaming"
	"math"
	"sort"
	"strings"
	"testing"
)

/*
func Test_Values(t *testing.T) {
	var vs []*Value = Values{&Value{val: 1}, &Value{val: 2}, &Value{val: 0}}
	stream := streaming.Of(vs)
	if err != nil {
		return
	}
	stream.Limit(5).ForEach(func(i interface{}) {
		fmt.Printf("%v\n", *i.(*Value))
	})

	fmt.Println()

	stream.MapSame(func(i interface{}) interface{} {
		return (*i.(*Value)).val * 100
	}).ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
	fmt.Printf("%v\n\n", stream.Count())
	fmt.Printf("%v\n\n", stream.Filter(func(i interface{}) bool {
		return i.(*Value).val > 1
	}).Count())

	stream.Limit(3).ForEach(func(i interface{}) {
		fmt.Printf("%v\t", i.(*Value).val)
	})
}*/

var ints = streaming.Ints{1, 5, 7, 2, 6, 9, 3}

func Test_Of(t *testing.T) {
	s := streaming.Of(ints)
	if err != nil {
		return
	}
	fmt.Printf("%v\n", s)
}

func Test_Filter(t *testing.T) {
	s := streaming.Of(ints)
	if err != nil {
		return
	}
	filter := s.Filter(func(i interface{}) bool {
		return i.(int) > 2
	})
	fmt.Printf("%v\n", filter)
}

func Test_Collect(t *testing.T) {
	s := streaming.Of(ints)
	if err != nil {
		return
	}
	collect := s.Filter(func(i interface{}) bool {
		return i.(int) > 2
	}).Collect()
	fmt.Printf("%v\n", collect)
}

func Test_Map(t *testing.T) {
	s := streaming.Of(ints)
	if err != nil {
		return
	}
	collect := s.MapSame(func(i interface{}) interface{} {
		return i.(int) * 2
	}).Collect()
	fmt.Printf("%v\n", collect)
}

func Test_ForEach(t *testing.T) {
	s := streaming.Of(ints)
	if err != nil {
		return
	}
	s.ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func Test_Limit(t *testing.T) {
	s := streaming.Of(ints)
	if err != nil {
		return
	}
	s.Limit(3).ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func Test_Array(t *testing.T) {
	var arr = ints
	s := streaming.Of(arr)
	if err != nil {
		return
	}
	s.Map(func(i interface{}) interface{} {
		return math.Pow(float64(i.(int)), 2)
	}).ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
	fmt.Println()
	println(s.Count())
}

func TestStream_Reduce(t *testing.T) {
	s := streaming.Of(ints)
	if err != nil {
		return
	}

	reduce := s.Reduce(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})
	fmt.Println(reduce)
}

func Test_nil(t *testing.T) {
	stream := streaming.Of(nil)
	if stream == nil {
		return
	}

	println(stream.MapSame(func(i interface{}) interface{} {
		return i.(int) * 100
	}).Filter(func(i interface{}) bool {
		return i.(int) > 3
	}).Limit(100).Count())
}

func Test_Distinct(t *testing.T) {
	stream := streaming.Of(ints)
	if err != nil {
		return
	}
	stream.Distinct().ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func Test_Sum(t *testing.T) {
	stream := streaming.Of(ints)
	sum := stream.Sum(func(i interface{}) float64 {
		return float64(i.(int))
	})
	fmt.Printf("%v\n", sum)
}

func Test_Match(t *testing.T) {
	stream := streaming.Of(ints)
	println(stream.AnyMatch(func(i interface{}) bool {
		return i.(int) > 12
	}))

	println(stream.AllMatch(func(i interface{}) bool {
		return i.(int) > 0
	}))

	println(stream.NoneMatch(func(i interface{}) bool {
		return i.(int) == 0
	}))
}

func Test_IsEmpty(t *testing.T) {
	println(streaming.Of(nil).IsEmpty())
}

func Test_FlatMap(t *testing.T) {
	stream := streaming.Of(words)
	flatMap := stream.FlatMap(func(i interface{}) streaming.Slicer {
		split := strings.Split(i.(string), " ")
		return streaming.Strings(split)
	})
	flatMap.ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

var words = streaming.Strings{"one", "two", "three"}

func Test_Peek(t *testing.T) {
	stream := streaming.Of(words)
	collect := stream.Peek(func(i interface{}) {
		fmt.Printf("%v is consumed\n", i)
	}).Collect()
	fmt.Printf("%v\n", collect)
}

func Test_Skip(t *testing.T) {
	stream := streaming.Of(ints)
	collect := stream.Skip(1).Collect()
	fmt.Printf("%v\n", collect)
}

func Test_FilterCount(t *testing.T) {
	stream := streaming.Of(ints)
	println(stream.FilterCount(func(i interface{}) bool {
		return i.(int) > 1
	}))
}

func Test_FindFirst(t *testing.T) {
	stream := streaming.Of(ints)
	first := stream.FindFirst()
	fmt.Printf("%v\n", first)
}

func Test_Copy(t *testing.T) {
	s := streaming.Of(ints)
	ss := s.Copy()
	fmt.Printf("%p %p\n", s, s.Collect())
	fmt.Printf("%p %p\n", ss, ss.Collect())
}

func Test_std_Sort(t *testing.T) {
	s := streaming.Of(ints)
	c := s.Collect()
	sort.Slice(c, func(i, j int) bool {
		return c[i].(int) > c[j].(int)
	})
	fmt.Printf("%v\n", c)

	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})
	fmt.Printf("%v\n", ints)
}

func Test_Stream_Sort(t *testing.T) {
	s := streaming.Of(ints)
	slice := s.Sorted(func(i, j int) bool {
		//return s.Element(i).(int) > s.Element(j).(int)
		return s.Element(i).(int) > s.Element(j).(int)
	}).Collect()
	fmt.Printf("%v\n", slice)
}

func Test_Stream_Copy_Sort(t *testing.T) {
	s := streaming.Of(ints)
	_copy := s.Copy()
	slice := s.Sorted(func(i, j int) bool {
		return s.Element(i).(int) > s.Element(j).(int)
	}).Collect()
	fmt.Printf("%v\n", slice)
	fmt.Printf("%v\n", _copy.Collect())
}

func Test_Element(t *testing.T) {
	s := streaming.Of(ints)
	fmt.Printf("%v\n", s.Element(1))
}
