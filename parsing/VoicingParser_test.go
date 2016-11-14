package parsing

import (
	"github.com/cptavatar/rofretta-go/theory"
	"testing"
)

func createTestParser() VoicingParser {
	return CreateVoicingParser(theory.CreateInstrument(theory.GuitarStandardTuning))
}

func TestParseVoicing(t *testing.T) {
	parser := createTestParser()
	voicing, valid := parser.ParseVoicing("xx0331 xxx341 Bb maj")
	if ! valid {
		t.Error("ParseVoicing return invalid for valid chord defintion")
	}
	if voicing.Root().Root().Name() != "B" || voicing.Root().Modifier() != -1 {
		t.Error("Parser should have identifed root as Bb")
	}
	if voicing.Chord().ShortName() != "maj" {
		t.Error("Parser should have identifed chord as major")
	}
	stringz := voicing.Strings()
	if (! stringz[0].Disabled() ) {
		t.Error("low E string should be muted")
	}
	if ( stringz[2].Fret() != 0 || stringz[2].Finger() != -1 || stringz[2].Disabled()) {
		t.Errorf("D string should be open, no finger: fret:%d, finger:%d, disabled:%t ", stringz[2].Fret(), stringz[2].Finger(), stringz[2].Disabled())
	}
	if ( stringz[5].Fret() != 1 || stringz[5].Finger() != theory.Index || stringz[5].Disabled()) {
		t.Errorf("high E string should be should be index, 1st fret: fret:%d, finger:%d, disabled:%t ", stringz[5].Fret(), stringz[5].Finger(), stringz[5].Disabled())
	}
}

