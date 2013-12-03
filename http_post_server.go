package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var buffer []byte

var maxBuffer []byte

func requestHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Panic(err)
	}
	buffer = data
	maxBuffer = append(maxBuffer, data...)
	maxBuffer = append(maxBuffer, "<br><br>"...)

	log.Printf("%s\n", data)

	fmt.Fprintf(w, "Ok")
}

// Reset logs
func requestReset(w http.ResponseWriter, r *http.Request) {
	maxBuffer = nil
}

// Show POST-request logs
func requestLog(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	fmt.Fprintf(w, "<html><body>%s</body></html>", maxBuffer)
}

func main() {
	http.HandleFunc(Config().Url, requestHandler)
	http.HandleFunc("/reset", requestReset)
	http.HandleFunc("/log", requestLog)
	http.ListenAndServe(Config().Port, nil)
}
