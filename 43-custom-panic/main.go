package main

func main() {

	fn := GetFullName("", "Palaparthi")
	println(fn)

}

func GetFullName(firstname, lastname string) string {
	if firstname == "" {
		panic("firstname cannot be empty")
	}
	if lastname == "" {
		panic("lastname cannot be empty")
	}

	return firstname + lastname
}
