package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestGetField(t *testing.T) {
	timeNow := time.Now()
	TestTrack := Track{Hdate: timeNow, Pilot: "Mariusz", Glider: "BestGlider", GliderID: "BestGliderID", TrackLength: 400.25}

	hdate, ok := TestTrack.getField("H_date")
	if !ok {
		t.Errorf("Coudln't find field H_date")
	}
	if hdate != timeNow.String() {
		t.Errorf("TestTrack.Hdate %s doesn't match expected value %s", hdate, timeNow.String())
	}

	pilot, ok := TestTrack.getField("pilot")
	if !ok {
		t.Errorf("Couldn't find field pilot")
	}
	if pilot != "Mariusz" {
		t.Errorf("TestTrack.Pilot %s doesn't match expected value Mariusz", pilot)
	}

	glider, ok := TestTrack.getField("glider")
	if !ok {
		t.Errorf("Couldn't find field glider")
	}
	if glider != "BestGlider" {
		t.Errorf("TestTrack.Glider %s doesn't match expected value BestGlider", glider)
	}

	gliderID, ok := TestTrack.getField("glider_id")
	if !ok {
		t.Errorf("Couldn't find field glider_id")
	}
	if gliderID != "BestGliderID" {
		t.Errorf("TestTrack.GliderID %s doesn't match expected value BestGliderID", gliderID)
	}

	trackLength, ok := TestTrack.getField("track_length")
	if !ok {
		t.Errorf("Coudln't find field track_length")
	}
	if trackLength != fmt.Sprintf("%f", 400.25) {
		t.Errorf("TestTrack.TrackLength %s doesn't match expected value %f", trackLength, 400.25)
	}

}

func TestAdd(t *testing.T) {
	timeNow := time.Now()
	TestTrack := Track{Hdate: timeNow, Pilot: "Mariusz", Glider: "BestGlider", GliderID: "BestGliderID", TrackLength: 400.25}
	TestTracks := Tracks{Tracks: make(map[uuid.UUID]Track)}

	tID, err := TestTracks.add(TestTrack)
	if err != nil {
		t.Errorf("Coudln't add TestTrack to TestTracks")
	}
	track, ok := TestTracks.Tracks[tID]
	if !ok {
		t.Errorf("Couldn't find track with id %s", tID.String())
	}
	if track != TestTrack {
		t.Errorf("TestTrack supplied to Tracks.add doesn't match the one within")
	}
}

func TestGetIDs(t *testing.T) {
	timeNow := time.Now()
	TestTrack1 := Track{Hdate: timeNow, Pilot: "Mariusz", Glider: "BestGlider", GliderID: "BestGliderID", TrackLength: 400.25}
	TestTrack2 := Track{Hdate: timeNow, Pilot: "Sindre", Glider: "SecondBestGlider", GliderID: "SecondBestGliderID", TrackLength: 395.19}
	TestTrackID1 := uuid.New()
	TestTrackID2 := uuid.New()
	TestTracks := Tracks{Tracks: make(map[uuid.UUID]Track)}
	TestTracks.Tracks[TestTrackID1] = TestTrack1
	TestTracks.Tracks[TestTrackID2] = TestTrack2

	IDs := TestTracks.getIDs()

	if len(IDs) != 2 {
		t.Errorf("Invalid length of TestTracks.getIDs, got %d, expected %d", len(IDs), 2)
	}
	if !((IDs[0] == TestTrackID1 && IDs[1] == TestTrackID2) || (IDs[1] == TestTrackID1 && IDs[0] == TestTrackID2)) {
		t.Errorf("Tracks.getIDs doesn't match manually added IDs")
	}
}

func TestGetTrack(t *testing.T) {
	timeNow := time.Now()
	TestTrack := Track{Hdate: timeNow, Pilot: "Mariusz", Glider: "BestGlider", GliderID: "BestGliderID", TrackLength: 400.25}
	TestTrackID := uuid.New()
	TestTracks := Tracks{Tracks: make(map[uuid.UUID]Track)}
	TestTracks.Tracks[TestTrackID] = TestTrack

	returnedTestTrack, ok := TestTracks.getTrack(TestTrackID)
	if !ok {
		t.Errorf("Coudln't get track using getTrack(%s)", TestTrackID)
	}
	if TestTrack != returnedTestTrack {
		t.Errorf("TestTrack added to TestTracks doesn't match the one returned")
	}
}
