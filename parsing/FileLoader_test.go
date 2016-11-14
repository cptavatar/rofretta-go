package parsing

import (
	"testing"
	"github.com/cptavatar/rofretta-go/theory"
)

func TestLoadVoicings(t *testing.T) {
	voicings := LoadVoicings("testdata", theory.CreateInstrument(theory.GuitarStandardTuning))
	if len(voicings) != 20 {
		t.Errorf("Expected there to be 20 voicings loaded, there were %d", len(voicings))
	}
	voicing := voicings[19]
	if voicing.Root().Natural().Name() != "B" || voicing.Root().Modifier() != 0 {
		t.Error("Loader should have picked up last voicing root as B")
	}
	if voicing.Chord().ShortName() != "maj" {
		t.Error("Loader should have picked up last voicing chord as major")
	}
}
