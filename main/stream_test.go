package main

import (
	"fmt"
	"github.com/gocurr/good/streaming"
	"math"
	"reflect"
	"sort"
	"testing"
)

type Value struct {
	val float64
}

type Values []*Value

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

	stream.Map(func(i interface{}) interface{} {
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
}

func TestWrap(t *testing.T) {
	s := streaming.Of([]int{1, 2, 3})
	if err != nil {
		return
	}
	fmt.Printf("%v\n", s)
}

func Test_Filter(t *testing.T) {
	s := streaming.Of([]int{1, 2, 3})
	if err != nil {
		return
	}
	filter := s.Filter(func(i interface{}) bool {
		return i.(int) > 2
	})
	fmt.Printf("%v\n", filter)
}

func Test_Collect(t *testing.T) {
	s := streaming.Of([]int{1, 2, 3})
	if err != nil {
		return
	}
	collect := s.Filter(func(i interface{}) bool {
		return i.(int) > 2
	}).Collect()
	fmt.Printf("%v\n", collect)
}

func Test_Map(t *testing.T) {
	s := streaming.Of([]int{1, 2, 3})
	if err != nil {
		return
	}
	collect := s.Map(func(i interface{}) interface{} {
		return i.(int) * 2
	}).Collect()
	fmt.Printf("%v\n", collect)
}

func Test_ForEach(t *testing.T) {
	s := streaming.Of([]int{1, 2, 3})
	if err != nil {
		return
	}
	s.ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func Test_Limit(t *testing.T) {
	s := streaming.Of([]int{1, 2, 3})
	if err != nil {
		return
	}
	s.Limit(3).ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func Test_Array(t *testing.T) {
	var arr = [...]int{1, 2, 3}
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
	s := streaming.Of([]int{11, 3})
	if err != nil {
		return
	}

	reduce, i := s.Reduce(func(a, b interface{}) bool {
		return a.(int) > b.(int)
	})
	fmt.Println(reduce, i)
}

func Test_nil(t *testing.T) {
	var raw []int
	stream := streaming.Of(raw)
	if err != nil {
		return
	}

	println(stream.Map(func(i interface{}) interface{} {
		return i.(int) * 100
	}).Filter(func(i interface{}) bool {
		return i.(int) > 3
	}).Limit(100).Count())
}

func Test_Distinct(t *testing.T) {
	//raw := []Value{{val: 1}, {val: 2}, {val: 2}, {val: 1}}
	v1 := &Value{1}
	v2 := &Value{2}
	raw := []*Value{v1, v2, v2}
	stream := streaming.Of(raw)
	if err != nil {
		return
	}
	stream.Distinct().ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func Test_Sum(t *testing.T) {
	stream := streaming.Of([]int{1, 2, 3})
	sum := stream.Sum(func(i interface{}) float64 {
		return float64(i.(int))
	})
	fmt.Printf("%v\n", sum)
}

func Test_Match(t *testing.T) {
	stream := streaming.Of([]int{1, 2, 3})
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
	stream := streaming.Of([]string{"hello there", "good morning"})
	flatMap := stream.FlatMap(func(i interface{}) interface{} {
		return [...]string{} //strings.Split(i.(string), " ")
	})
	flatMap.ForEach(func(i interface{}) {
		fmt.Printf("%v\n", i)
	})
}

func Test_Peek(t *testing.T) {
	stream := streaming.Of([]string{"one", "two", "three"})
	collect := stream.Peek(func(i interface{}) {
		fmt.Printf("%v is consumed\n", i)
	}).Collect()
	fmt.Printf("%v\n", collect)
}

func Test_Skip(t *testing.T) {
	stream := streaming.Of([]int{1, 2, 3})
	collect := stream.Skip(3).Collect()
	fmt.Printf("%v\n", collect)
}

func Test_FilterCount(t *testing.T) {
	stream := streaming.Of([]int{1, 2, 3})
	println(stream.FilterCount(func(i interface{}) bool {
		return i.(int) > 1
	}))
}

func Test_FindFirst(t *testing.T) {
	stream := streaming.Of([]int{2, 1, 3})
	first := stream.FindFirst()
	fmt.Printf("%v\n", first)
}

func Test_FlatMapX(t *testing.T) {
	var a = [...]int{1, 5}
	var b = []int{2, 3}
	var raw = []interface{}{b, a}
	stream := streaming.Of(raw)
	slice := stream.FlatMap(func(i interface{}) interface{} {
		switch reflect.TypeOf(i).Kind() {
		case reflect.Int:
			return i.(int) * 2
		case reflect.Slice, reflect.Array:
			return i
		}
		return nil
	}).Collect()
	fmt.Printf("%v\n", slice)
}

func Test_Copy(t *testing.T) {
	s := streaming.Of([]int{1, 2, 0})
	ss := s.Copy()
	fmt.Printf("%p %p\n", s, s.Collect())
	fmt.Printf("%p %p\n", ss, ss.Collect())
}

func Test_std_Sort(t *testing.T) {
	s := streaming.Of([]int{1, 12, 9})
	c := s.Collect()
	sort.Slice(c, func(i, j int) bool {
		return c[i].(int) > c[j].(int)
	})
	fmt.Printf("%v\n", c)

	ints := []int{1, 5, 3}
	sort.Slice(ints, func(i, j int) bool {
		return ints[i] < ints[j]
	})
	fmt.Printf("%v\n", ints)
}

func Test_Stream_Sort(t *testing.T) {
	s := streaming.Of([]int{1, 3, 2, 9, 0, 5, 4, 6, 8, 7})
	slice := s.Sort(func(i, j int) bool {
		//return s.Element(i).(int) > s.Element(j).(int)
		return s.Element(i).(int) < s.Element(j).(int)
	}).Collect()
	fmt.Printf("%v\n", slice)
}

func Test_Stream_Copy_Sort(t *testing.T) {
	s := streaming.Of([]int{1, 3, 2, 9, 0, 5, 4, 6, 8, 7})
	ss := s.Copy()
	slice := ss.Sort(func(i, j int) bool {
		return ss.Element(i).(int) < ss.Element(j).(int)
	}).Collect()
	fmt.Printf("%v\n", slice)
	fmt.Printf("%v\n", s.Collect())
}
