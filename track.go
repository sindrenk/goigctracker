package main

import (
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

func (ts *Tracks) getIDs() []uuid.UUID {
	ids := make([]uuid.UUID, len(ts.Tracks))
	i := 0
	for k := range ts.Tracks {
		ids[i] = k
	}

	return ids
}
