package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

var startTime time.Time
var tracks Tracks

func main() {
	startTime = time.Now()
	tracks.init()

	http.HandleFunc("/igcinfo/api/", rootHandler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
