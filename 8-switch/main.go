package main

import "fmt"

func main() {

	var day uint8 = 4

	switch day {
	case 1:
		println("Sunday")
	case 2:
		println("Monday")
	case 3:
		println("Tuesday")
	case 4:
		println("Wednesday")
	case 5:
		println("Thursday")
	case 6:
		println("Friday")
	case 7:
		println("Saturday")
	default:
		println("noday")
	}

	num := 32
	// empty switch
	switch {
	case num%8 == 0:
		println(num, "is divisible by 8")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%2 == 0:
		println(num, "is divisible by 2")
	}

	// fallthrough
	num = 36
	// empty switch
	switch {
	case num%2 == 0:
		println(num, "is divisible by 2")
		fallthrough
	case num%4 == 0:
		println(num, "is divisible by 4")
		fallthrough
	case num%8 == 0:
		println(num, "is divisible by 8")
	}

	char := 'a'

	switch char {
	case 'A', 'E', 'I', 'O', 'U':
		fmt.Println(string(char), " is a upper case vovel")
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println(string(char), " is a lower case vovel")
	default:
		fmt.Println(string(char), " is a consonent or a special unicode char")
	}

}
