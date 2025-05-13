package main

import "fmt"

func main() {

	slice := []int{1, 2, 3}

	SqOfElems1(&slice, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(slice)
	fmt.Printf("%p\n", &slice[0])
	fmt.Printf("%p\n", &slice)
}

func SqOfElems1(slice *[]int, nums ...int) {
	fmt.Printf("%p\n", slice)
	if slice != nil {
		*slice = append(*slice, nums...) //
		for i, v := range *slice {
			(*slice)[i] = v * v
		}
	}
	// ptr: 0x11 len:4 cap:5
}
