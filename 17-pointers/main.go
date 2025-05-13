package main

func main() {
	var num1 int = 100
	var ptr1 *int = &num1

	var ptr4 *uint8 // 8 bytes

	if ptr4 != nil {
		*ptr4 = 100
	}

	if ptr1 == nil {
		println("nil pointer")
	}
	ptr5 := new(int) // new function allocated some space for int
	IncrP(ptr5)
	println(*ptr5)
	{
		ptr6 := SqH(*ptr5)
		println(*ptr6)
	}

	ptr7 := new(int)

	SqS(20, ptr7)
	println(*ptr7)

	Incr(num1)
	println(num1)
	IncrP(ptr1)
	IncrP(&num1)
	println(num1)

	slice1 := make([]int, 10)
	println(&slice1)

	slice2 := make([]int, 10000)
	println(&slice2)

	var arr1 [10]int
	println(&arr1[0])
	//fmt.Println(arr1) // Dont use fmt.Println
	for i := 0; i < len(arr1); i++ {
		print(arr1[i])
	}
	println()

	var arr2 [100000]int
	println(&arr2[0])

}

func IncrP(v *int) {
	if v != nil {
		*v++
	}
}

func Incr(v int) {
	v++
}

// 1. Deref a nil pointer --> Yes problem in go as well
// 2. Double Free --> since you are not going to free/delete/drop no pointer of double free
// 3. Dangling Pointer --> There is a problem but Go solves it
// 4. Memory Leak --> Since GC, memory is deallocated

func SqH(num int) *int {
	ptr := new(int)  // create memory
	*ptr = num * num // derefe
	return ptr       // returns the address //dangling pointer
}

func SqS(num int, ptr *int) {
	if ptr != nil {
		*ptr = num * num // derefe
	}
}
