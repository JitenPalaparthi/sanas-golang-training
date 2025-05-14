package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Hello World")
	fw1 := NewFileWriter("data.txt")
	fmt.Fprintln(fw1, "Hello World")

	var fw2 *FileWriter

	_, err := fmt.Fprintln(fw2, "Hello Sanas.io")

	if err != nil {
		filerror := err.(*FileError)
		fmt.Println("Code:", filerror.Code)
		fmt.Println("Msg:", filerror.Msg)
	}

}

type FileWriter struct {
	FileName string
}

type FileError struct {
	Code, Msg string
}

func NewFileError(code, msg string) *FileError {
	return &FileError{code, msg}
}

func (fe *FileError) Error() string {
	return fmt.Sprint("Code:", fe.Code, "Msg:", fe.Msg)
}

func (fw *FileWriter) Write(p []byte) (n int, err error) {
	if fw == nil {
		return 0, NewFileError("100", "nil file writer")
	}
	if fw.FileName == "" {
		return 0, NewFileError("101", "invalid file name")
	}

	f, err := os.OpenFile(fw.FileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(p)
}

// type Rect struct {
// 	L, B float32
// }

//	func NewRect(l, b float32) *Rect {
//		return &Rect{l, b}
//	}
func NewFileWriter(fname string) *FileWriter {
	return &FileWriter{fname}
}

func DefaultFileWriter() *FileWriter {
	return &FileWriter{"data.txt"}
}
