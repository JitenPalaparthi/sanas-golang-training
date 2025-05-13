package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {

	var any1 any // every type is implenenting this

	any1 = "100"

	switch any1.(type) { // type switch
	case int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
		fmt.Println(any1, "is a number", reflect.TypeOf(any1))
	default:
		fmt.Println(any1, "is not a number and the type is", reflect.TypeOf(any1))
	}
	fmt.Println(IsNumber(nil))

	var a, b any
	a, b = 100, "Hello World"

	if s, err := add(a, b); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("sum is %.3f\n", s)
	}

	a, b = 100, 100

	if s, err := add(a, b); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("sum is %.3f\n", s)
	}

	a, b = 100.12, 100

	if s, err := add(a, b); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("sum is %.3f\n", s)
	}

	a, b = true, false
	if s, err := add(a, b); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("sum is %.3f\n", s)
	}

	a, b = int8(100), int8(20)
	if s, err := add(a, b); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("sum is %.3f\n", s)
	}

	a, b = 100.12, 100.45

	if s, err := add(a, b); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("sum is %.3f\n", s)
	}

	a, b = 100.1234, float32(100.45)

	if s, err := add(a, b); err != nil {
		println(err.Error())
	} else {
		fmt.Printf("sum is %.3f\n", s)
	}

	s := sum(100, 200)
	fmt.Printf("sum is %.3f\n", s)

	s = sum(100.123, 200.123)
	fmt.Printf("sum is %.3f\n", s)

	s = sum(float32(100.123), float32(200.123))
	fmt.Printf("sum is %.3f\n", s)

	s, err := add(any(100), any(false))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("sum is %.3f\n", s)
	}

}

// type earser vs monomorphization

func IsNumber(a any) (bool, error) {
	switch a.(type) { // type switch
	case int, uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
		return true, nil
	case nil:
		return false, errors.New("input is nil")
		//return false, fmt.Errorf("input is nil")
	default:
		return false, nil
	}
}

func add(a, b any) (sum float64, err error) {
	//sum := float64(0)

	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		return 0, fmt.Errorf("input arguments are not of same type a:%T b:%T", a, b)
	}

	if ok, _ := IsNumber(a); !ok {
		return 0, errors.New("first argument a is not a number")
	}

	switch a1 := a.(type) {
	case int:
		sum = float64(a1 + b.(int))
	case uint:
		sum = float64(a1 + b.(uint))
	case int8:
		sum = float64(a1 + b.(int8))
	case int16:
		sum = float64(a1 + b.(int16))
	case int32:
		sum = float64(a1 + b.(int32))
	case int64:
		sum = float64(a1 + b.(int64))
	case uint8:
		sum = float64(a1 + b.(uint8))
	case uint16:
		sum = float64(a1 + b.(uint16))
	case uint32:
		sum = float64(a1 + b.(uint32))
	case uint64:
		sum = float64(a1 + b.(uint64))
	case float32:
		sum = float64(a1 + b.(float32))
	case float64:
		sum = a1 + b.(float64)
	}

	// if ok, _ := IsNumber(b); !ok {
	// 	return 0, errors.New("second argument b is not a number")
	// }

	return sum, nil
}

func sum[T int | uint | int8 | int16 | int32 | int64 | uint8 | uint16 | uint32 | uint64 | float32 | float64](a T, b T) float64 {
	return float64(a + b)
}
