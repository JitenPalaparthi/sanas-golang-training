package main

import "fmt"

func main() {

	slice1 := make([]int, 2)
	slice1[0], slice1[1] = 2, 4
	//slice1 = SqOfElems(slice1, 6, 8) // ptr:0x11 len: 2 cap:5
	SqOfElems1(slice1, 6, 8)
	fmt.Println(slice1)

	slice2 := make([]int, 2)
	copy(slice2, slice1)

	slice3 := make([]int, 10)
	copy(slice3, slice2)

	fmt.Println(slice3)

	slice := make([]int, 10, 15)
	for i := 0; i < 10; i++ {
		slice[i] = i + 1
	}

	// Ptr:0x11 Len:10 Cap:10

	slice4 := slice[:]   // slice4 := slice Len 10 Cap 15
	slice5 := slice[5:]  // slice4 := slice Len 5 Cap
	slice6 := slice[:5]  // slice4 := slice
	slice7 := slice[3:8] // slice4 := slice
	PrintSliceHeader("Slice4", slice4)
	PrintSliceHeader("Slice5", slice5)
	PrintSliceHeader("Slice6", slice6)
	PrintSliceHeader("Slice7", slice7)
	// 1 2 3 4 5 6 7 8 9 10
	slice8 := append(slice[:5], slice[6:]...)
	fmt.Println(slice8)

	clear(slice8)
	fmt.Println(slice8)
	slice8 = nil
	if slice8 == nil {
		println("Yes nil")
	}
}

func SqOfElems1(slice []int, nums ...int) {
	slice = append(slice, nums...) //
	for i, v := range slice {
		slice[i] = v * v
	}
	// ptr: 0x11 len:4 cap:5
}
func SqOfElems2(slice []int, nums ...int) []int {
	slice = append(slice, nums...) //
	for i, v := range slice {
		slice[i] = v * v
	}
	return slice
	// ptr: 0x11 len:4 cap:5
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
