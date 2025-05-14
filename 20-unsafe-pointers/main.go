package main

import (
	"fmt"
	"unsafe"
)

func main() {

	arr1 := [5]int{100, 200, 300, 400, 500}

	// pointer arthimetic in Go (not directly)

	//ptr := &arr1[0]
	// ptr += 8 // arithmetic operatiion on a pointer
	// value or dereference a pointer

	uintptr1 := uintptr(unsafe.Pointer(&arr1[0])) // 1 and 4

	uintptr1 += unsafe.Sizeof(arr1[0]) // arithmetic operation

	ptr := (*int)(unsafe.Pointer(uintptr1)) // 3 and 2
	println(*ptr)

	//var uintptr2 uintptr
	for i := 0; i < len(arr1); i++ {
		uintptr1 := uintptr(unsafe.Pointer(&arr1[i])) // 1 and 4
		fmt.Printf("%p--->0x%x\n", &arr1[i], uintptr1)
		ptr := *(*int)(unsafe.Pointer(uintptr1)) // 3 and 2
		uintptr1 += unsafe.Sizeof(arr1[i])       // arithmetic operation
		print(ptr, " ")
		//uintptr2 = uintptr1
	}
	println()

	// uintptr2 += 36
	// ptr3 := *(*int)(unsafe.Pointer(uintptr2)) // 3 and 2\
	// print(ptr3, " ")

	str1 := "Hello World!"
	str2 := "Hello Fake string!"

	// Ptr --> 8
	// Len --> 8

	arr2 := (*[2]int)(unsafe.Pointer(&str1))

	arr3 := (*[2]int)(unsafe.Pointer(&str2))

	println(arr2[0]) // derefernece
	println((*arr2)[1])

	(*arr2)[1] = 16
	(*arr2)[0] = arr3[0]
	println(str1, len(str1))

	// *void --> unsafe.Pointer

}

// 1. A pointer value of any type can be converted to a Pointer.
// 2. A Pointer can be converted to a pointer value of any type.
// 3. A uintptr can be converted to a Pointer.
// 4. A Pointer can be converted to a uintptr.
