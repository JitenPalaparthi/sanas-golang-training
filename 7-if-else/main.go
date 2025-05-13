package main

import "strconv"

func main() {

	if true || true && false {
		println("true")
	} else {
		println("false")
	}

	var str string = "12312"

	if v, err := strconv.Atoi(str); err != nil {
		println(err.Error())
	} else if num := 2; v%num == 0 {
		println("even number", v)
	} else {
		println("odd number", v)
	}

	str = "12312.123"

	if v1, err1 := strconv.ParseFloat(str, 32); err1 != nil {
		println(err1.Error())
	} else {
		println(v1)
	}

	str = `} else if v%2 == 0 {`

}
