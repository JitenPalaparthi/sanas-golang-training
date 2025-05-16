package main

import (
	"demo/handlers"
	"demo/utils"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

var (
	PORT     string
	FileName string
)

func main() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		flag.StringVar(&PORT, "port", "8080", "--port=8080 or -port=8080 or --port 8080")

	}
	if FileName == "" {
		flag.StringVar(&FileName, "filename", "data.txt", "--filename=data.txt")
	}
	flag.Parse()
	utils.Init(FileName)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Sanas.ai")
		w.Write([]byte("Hello Sanas.ai"))
		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/ping", Ping)
	http.HandleFunc("/user", handlers.CreateUser)
	http.HandleFunc("/user1", handlers.CreateUser1)

	log.Println("application started and running on port:", PORT)
	if err := http.ListenAndServe(":"+PORT, nil); err != nil {
		log.Println(err.Error())
		runtime.Goexit() // wait for all other goroutines do the work
	} // 0.0.0.0 192.168.0.1 127.0.0.1
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong"))
	w.WriteHeader(http.StatusOK)
}
