package handlers

import (
	"demo/models"
	"demo/utils"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			w.Write([]byte("opps something went"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user := new(models.User)
		err = json.Unmarshal(b, user)
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte("some issue"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = user.Validate()
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.Status = "active"
		user.LastModified = uint64(time.Now().Unix())
		utils.ChUser <- *user

		err = <-utils.HasFileStored
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.Write([]byte("User successfully created"))
		w.WriteHeader(201)

	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}

}

func CreateUser1(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		b, err := io.ReadAll(r.Body)

		if err != nil {
			log.Println(err.Error())
			w.Write([]byte("opps something went"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		user := new(models.User)
		err = json.Unmarshal(b, user)
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte("some issue"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = user.Validate()
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		user.Status = "active"
		user.LastModified = uint64(time.Now().Unix())
		utils.ChUser <- *user
		chErr := make(chan error)

		go func(filename string, cherr chan error) {
			f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
			if err != nil {
				cherr <- err
			}
			defer f.Close()

			_, err = f.Write([]byte(user.ToString()))
			if err != nil {
				cherr <- err
			} else {
				cherr <- nil
			}

		}("data2.txt", chErr)

		err = <-chErr

		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Write([]byte("User successfully created"))
		w.WriteHeader(201)

	} else {
		w.WriteHeader(http.StatusNotImplemented)
	}

}
