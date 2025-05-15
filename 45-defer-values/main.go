package main

func main() {
	str := "Hello World"
	defer println()
	for _, v := range str {
		defer print(string(v))
	}
	num1 := 1

	defer func() {
		println("defer -- num1 global:", num1)
	}()

	num1 += 1
	println(num1)

	num2 := 1

	defer func(num int) {
		println("defer -- num:", num)
	}(num2)

	num2 += 1
	println(num2)

	num3 := 1
	defer println("defer -- num directly:", num3)
	num3 += 1
	println(num3)

}
