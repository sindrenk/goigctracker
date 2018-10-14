package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/marni/goigc"
)

// Metainfo is the struct that's replied on /igcinfo/api/
type Metainfo struct {
	Uptime  string `json:"uptime"`
	Info    string `json:"info"`
	Version string `json:"version"`
}

// Handles all requests to /igcinfo/api/ and calls on the appropriate function
func rootHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	if r.URL.Path == "/igcinfo/api/" {
		infoHandler(w, r)
		return
	} else if r.URL.Path == "/igcinfo/api/igc/" {
		igcHandler(w, r)
		return
	} else if id, err := uuid.Parse(parts[4]); strings.HasPrefix(r.URL.Path, "/igcinfo/api/igc/") && err == nil && len(parts) < 6 {
		trackHandler(w, r, id)
		return
	} else if id, err := uuid.Parse(parts[4]); strings.HasPrefix(r.URL.Path, "/igcinfo/api/igc/") && err == nil && len(parts[5]) > 0 {
		trackFieldHandler(w, r, id, parts[5])
		return
	}

	http.NotFound(w, r)
}

// Gets called on by rootHandler, serves JSON for metadata about the API
func infoHandler(w http.ResponseWriter, r *http.Request) {
	info := Metainfo{Uptime: timeSince(startTime), Info: "Service for IGC tracks.", Version: "v1"}
	js, err := json.Marshal(info)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// Gets called on my rootHandler, on POST it recives an url of a new IGC file to save, on GET it serves an array of the ID's of the saved tracks
func igcHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		decoder := json.NewDecoder(r.Body)

		// to map from "url": to "<url>"
		url := make(map[string]string)
		err := decoder.Decode(&url)
		// coulnd't parse POST data
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		track, err := igc.ParseLocation(url["url"])
		// coudln't get track from url, probably a bad URL in POST request
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		id, err := tracks.add(Track{Hdate: track.Date, Pilot: track.Pilot, Glider: track.GliderType, GliderID: track.GliderID, TrackLength: track.Points[0].Distance(track.Points[len(track.Points)-1])})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(id)
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tracks.getIDs())
	}
}

// Gets called on by rootHandler, recives the ID of a track then returns its data as JSON
func trackHandler(w http.ResponseWriter, r *http.Request, id uuid.UUID) {
	track, ok := tracks.getTrack(id)

	if !ok {
		http.NotFound(w, r)
		return
	}

	js, err := json.Marshal(track)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Gets called on by rootHandler, recives the ID of a track and the name of a field, provides that specific field as plain text
func trackFieldHandler(w http.ResponseWriter, r *http.Request, id uuid.UUID, field string) {
	track, ok := tracks.getTrack(id)

	if !ok {
		http.NotFound(w, r)
		return
	}

	f, ok := track.getField(field)

	if !ok {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, f)

}

// Returns a string with the time since time a in ISO 8601 format.
// Accounts for variying amount of days in a month.
// Based on https://play.golang.org/p/NgNnBW6gpq
func timeSince(a time.Time) string {
	b := time.Now()

	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year := int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)
	hour := int(h2 - h1)
	min := int(m2 - m1)
	sec := int(s2 - s1)

	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return fmt.Sprintf("P%dY%dM%dDT%dH%dM%dS", year, month, day, hour, min, sec)
}
