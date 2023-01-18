package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello!")
		d, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(rw, "Oops!", http.StatusBadRequest)
			// rw.WriteHeader(http.StatusBadRequest)
			// rw.Write([]byte("Oops!"))
			return
		}

		// log.Printf("Data %s\n", d)

		fmt.Fprintf(rw, "Hello %s", d)
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye!")
	})

	http.ListenAndServe(":9090", nil)
}
