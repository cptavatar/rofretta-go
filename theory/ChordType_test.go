package theory

import (
	"testing"
)

func TestChord_Notes(t *testing.T) {
	chord := CreateChord(CTMajor)
	c := CreateNote("C")

	notes := chord.Notes(c)
	if len(notes) != 3 {
		t.Error("C Major should have 3 notes")
	}
	if notes[0].Name() != "C" {
		t.Errorf("First note of C Maj should be C, was %s", notes[0].Name())
	}
	if notes[1].Name() != "E" {
		t.Errorf("Second note of C Maj should be E, was %s", notes[1].Name())
	}
	if notes[2].Name() != "G" {
		t.Errorf("Third note of C Maj should be G, was %s", notes[2].Name())
	}
}

func TestCreateChordsByIntervalCount(t *testing.T) {
	chords := CreateChordsByIntervalCount(3)
	if len(chords) == 0 {
		t.Error("We should have several chords for interval count 3, found none")
	}
	//for _,chord := range chords {
	//	fmt.Println(chord.Name())
	//}
}