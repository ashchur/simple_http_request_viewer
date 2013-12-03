package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func requestHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("%s\n", data)

	fmt.Fprintf(w, "Ok")
}

func main() {
	http.HandleFunc(Config().Url, requestHandler)
	http.ListenAndServe(Config().Port, nil)
}
