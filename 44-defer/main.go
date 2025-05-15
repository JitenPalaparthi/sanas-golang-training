package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("End of main")
	Write("data.log", []byte("Hello World"))
	defer func() { // func-1
		defer fmt.Println("End of defer func-1")
		panic("some panic") // the code panics.. 100%
		println("Hello Sanas")
		//
	}()
	fmt.Println("Start of main")
}

func Write(filename string, b []byte) (int, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
	n, err := f.Write(b)
	if err != nil {
		return 0, err
	}

	return n, nil
}
