package detection

import (
	"github.com/cptavatar/rofretta-go/theory"
)

type noteSetSizeDetector struct{}


// What we are going to do is attempt to identify the chord by the number of distinct
// notes it has in the voicing. We will take these distinct notes and map them to set of
// from 0-11 to deal with eharmonic equivilents. Next we find known chords that are
// made of the same number of distinct notes.
// Finally we brute force our way through the reduced set of candidates chords,
// trying each note as the root, building the notes for the candidate chord, and seeing if it
// matches the same tones in our enharmonic set
// BUG This algorithm doesn't work when notes are excluded, could just add a fallback for trying
// increasingly larger sizes if a exact match is not made (along with reverse check), or more tediously, define common
// dropped intervals / additional patterns
func (ns noteSetSizeDetector) Detect(vc theory.UnidentifedVoicing) []theory.Voicing {
	distinctNotes := make(map[string]theory.Note)

	for _, note := range vc.Notes(false) {
		distinctNotes[note.Name()] = note
	}
	size := len(distinctNotes)
	enharmonicNotes := make(map[int8]bool)
	for _, note := range distinctNotes {
		enharmonicNotes[note.EnharmonicValue()] = true
	}
	candidates := theory.CreateChordsByIntervalCount(size)

	retval := make([]theory.Voicing, 0, 10)
	if len(candidates) > 0 {
		for _, chord := range candidates {
			for _, possibleRoot := range distinctNotes {
				notes := chord.Notes(possibleRoot)
				if ns.checkCandidate(&enharmonicNotes, &notes) {
					retval = append(retval,
						theory.CreateVoicing(vc.Strings(), vc.Instrument(), possibleRoot, chord))
				}
			}
		}
	}

	return retval
}

func (ns noteSetSizeDetector) checkCandidate(desired *map[int8]bool, candidate *[]theory.Note) bool {
	for _, note := range (*candidate) {
		if !(*desired)[note.EnharmonicValue()] {
			return false
		}
	}
	return true
}