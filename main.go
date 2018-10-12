package main

import (
	"log"
	"net/http"
	"time"
)

var startTime time.Time
var tracks Tracks = Tracks{}

func main() {
	startTime = time.Now()
	tracks.init()

	http.HandleFunc("/igcinfo/api/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
