package main

import "fmt"

type SliceType[T string | int32 | float64] []T
type SliceMap[T string | int32] map[string]SliceType[T]

type NumberType[T int32 | float32] []T

type StructType[T int | string] struct {
	name   string
	salary T
}

func paramBack(value interface{}) interface{} {
	return value
}

func (s NumberType[T]) add() T {
	var rlt T
	for _, v := range s {
		rlt += v
	}

	return rlt
}

func Sum[T int32 | float32](list NumberType[T]) T {
	var rlt T
	for _, v := range list {
		rlt += v
	}

	return rlt
}

type Cube[T int32 | float32] struct {
	length T
}

type Cuboid[T int32 | float32] struct {
	width  T
	length T
	height T
}

func (c Cube[T]) calcVolum() T {
	return c.length * c.length * c.length
}

func (c Cuboid[T]) calcVolum() T {
	return c.length * c.width * c.height
}

func calcVolum[T int32 | float32](thing interface{}) T {
	cube, ok := thing.(Cube[T])

	if ok {
		return cube.calcVolum()
	} else {
		return 0
	}
}

func main() {
	fmt.Println(paramBack(123))
	fmt.Println(paramBack(123.33))
	fmt.Println(paramBack("123aaa"))

	var sliceInt NumberType[int32]
	sliceInt = append(sliceInt, 1)
	sliceInt = append(sliceInt, 2)
	sliceInt = append(sliceInt, 3)

	var sliceFloat NumberType[float32]

	sliceFloat = append(sliceFloat, 1.1)
	sliceFloat = append(sliceFloat, 1.2)
	sliceFloat = append(sliceFloat, 1.3)

	fmt.Println(sliceInt.add())
	fmt.Println(sliceFloat.add())

	res := Sum(sliceInt)

	fmt.Println(res)

	cube := Cube[int32]{
		length: 3,
	}

	volumn := calcVolum[int32](cube)

	fmt.Println(volumn)

}
