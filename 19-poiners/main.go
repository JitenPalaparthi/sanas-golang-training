package main

import "fmt"

func main() {

	str1 := "🚀 Hello World!"

	println(&[]byte(str1)[0])
	// &[]byte(str1)[0] &[]byte(str1)[1] &[]byte(str1)[2] &[]byte(str1)[3]

	slice1 := []int{10, 11, 12, 13, 14}
	slice1 = SqOfElems1(slice1, 15, 16, 17)
	fmt.Println(slice1)

	slice2 := []int{10, 11, 12, 13, 14}
	SqOfElems2(&slice2, 15, 16, 17)
	fmt.Println(slice2)

	num := 100
	ptr1 := &num
	ptr2 := &ptr1
	ptr3 := &ptr2

	***ptr3 = 2000
	println(num)

	//var any1 any = num
	//ptr4 := &any1

	a, b, c := 10, 20, 30
	slice3 := []*int{&a, &b, &c, nil}
	arr1 := [3]*int{&a, &b, &c}

	for _, v := range slice3 {
		if v != nil {
			*v = *v * 10
		}
	}

	println(a, b, c)

	d := 400

	for i, v := range arr1 {
		if v != nil {
			*v = *v * 10
		}
		if i == 2 {
			arr1[i] = &d
		}
	}

}

func SqOfElems1(slice []int, nums ...int) []int {
	fmt.Printf("%p\n", slice)
	if slice != nil {
		slice = append(slice, nums...) //
		for i, v := range slice {
			(slice)[i] = v * v
		}
	}
	return slice
	// ptr: 0x11 len:4 cap:5
}

func SqOfElems2(slice *[]int, nums ...int) {
	if slice != nil {
		fmt.Printf("%p\n", slice)
		if slice != nil {
			*slice = append(*slice, nums...) //
			for i, v := range *slice {
				(*slice)[i] = v * v
			}
		}
	}
	// ptr: 0x11 len:4 cap:5
}
