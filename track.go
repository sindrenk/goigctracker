package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// Track is the struct that contains information about a track
type Track struct {
	Hdate       time.Time `json:"H_date"`
	Pilot       string    `json:"pilot"`
	Glider      string    `json:"glider"`
	GliderID    string    `json:"glider_id"`
	TrackLength float64   `json:"track_length"`
}

// Tracks contains all Track-objects in memory
type Tracks struct {
	Tracks map[uuid.UUID]Track
}

// Returns the contents of a specific field, as is named in JSON
func (t *Track) getField(field string) (string, bool) {
	var f string
	ok := true
	switch field {
	case "H_date":
		f = t.Hdate.String()
	case "pilot":
		f = t.Pilot
	case "glider":
		f = t.Glider
	case "glider_id":
		f = t.GliderID
	case "track_length":
		f = fmt.Sprintf("%f", t.TrackLength)
	default:
		f = "INVALID FIELD"
		ok = false
	}

	return f, ok
}

// Makes the tracks.Tracks map. Must be called before struct Tracks is used.
func (ts *Tracks) init() {
	ts.Tracks = make(map[uuid.UUID]Track)
}

// Adds a track into a Tracks object
func (ts *Tracks) add(t Track) (uuid.UUID, error) {
	var id uuid.UUID
	var err error
	exists := true

	for exists {
		id, err = uuid.NewUUID()

		if err != nil {
			return uuid.Nil, err
		}

		if _, ok := ts.Tracks[id]; !ok {
			exists = false
			ts.Tracks[id] = t
		}
	}

	return id, nil
}

// Gets the IDs of all the saved tracks
func (ts *Tracks) getIDs() []uuid.UUID {
	ids := make([]uuid.UUID, len(ts.Tracks))
	i := 0
	for k := range ts.Tracks {
		ids[i] = k
		i++
	}

	return ids
}

// Gets a track from tracks by its ID
func (ts *Tracks) getTrack(id uuid.UUID) (Track, bool) {
	track, ok := ts.Tracks[id]

	return track, ok
}
