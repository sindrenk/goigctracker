package main

import (
	"log"
	"net/http"
	"time"
)

var startTime time.Time
var tracks Tracks

func main() {
	startTime = time.Now()
	tracks.init()

	http.HandleFunc("/igcinfo/api/", rootHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
