package main

import "fmt"

func main() {
	slice1 := []int{10, 11, 12, 13, 14}
	PrintSliceHeader("slice1 before append", slice1)
	slice1 = append(slice1, 15)
	PrintSliceHeader("slice1 after append", slice1)
	slice2 := slice1 // what heppen?
	slice2[0] = 99999
	fmt.Println(slice1)

	slice2 = append(slice2, 16)
	fmt.Println(slice1)

	slice2 = append(slice2, 17, 18, 19, 20, 21)
	slice2[1] = 88888
	fmt.Println(slice1)

	arr1 := [5]int{10, 20, 30, 40, 50}
	arr2 := [7]int{10, 20, 30, 40, 50, 60, 70}
	fmt.Println(SumOf(slice1))
	fmt.Println(SumOf(arr1[:]))
	fmt.Println(SumOf(arr2[:]))

	slice4 := arr1[:]
	slice4[0] = 99999
	fmt.Println(arr1)
	fmt.Printf("Addres of arr:%p\n", &arr1[0])
	fmt.Printf("Addres of slice4:%p\n", &slice4[0])
}

func PrintSliceHeader(name string, slice []int) {
	fmt.Println(name, "-------------------------")
	fmt.Printf("Address of slice:%p\n", &slice)
	if slice != nil && len(slice) > 0 {
		fmt.Println("Ptr in the slice:", &slice[0])
	} else {
		fmt.Println("Ptr in the slice:nil")
	}
	fmt.Println("Len:", len(slice))
	fmt.Println("Cap:", cap(slice))
	fmt.Println(slice)
	fmt.Println("-----------------------")
}

func SumOf(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SumOfElems(slice []int, nums ...int) int {
	slice = append(slice, nums...)
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func SqOfElems(slice []int, nums ...int) {
	slice = append(slice, nums...)
	for i, v := range slice {
		slice[i] = v * v
	}
}
