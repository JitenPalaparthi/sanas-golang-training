package utils

import (
	"demo/models"
	"log"
	"os"
)

var ChUser chan models.User
var HasFileStored chan error

//var FileName string

func Init(fileName string) {
	ChUser = make(chan models.User)
	go Process(fileName)
}

func Process(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	for user := range ChUser {
		_, err := f.Write([]byte(user.ToString()))
		if err != nil {
			HasFileStored <- err
		}
		log.Println(err)
	}
}
