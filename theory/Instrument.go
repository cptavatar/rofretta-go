package theory

type Instrument interface {
	Name() string
	MaxFret() uint8
	MidiNoteNumbers() []uint8
}

type InstrumentType int

const (
	GuitarStandardTuning InstrumentType = iota
)

type instrument struct {
	name            string
	maxFret         uint8
	midiNoteNumbers []uint8
}

func (instrument instrument ) Name() string {
	return instrument.name
}
func (instrument instrument ) MaxFret() uint8 {
	return instrument.maxFret
}
func (instrument instrument ) MidiNoteNumbers() []uint8 {
	return instrument.midiNoteNumbers
}

var guitarStandardTuning = instrument{
	name: "Guitar, Standard Tuning",
	maxFret: 24,
	midiNoteNumbers: []uint8{40, 45, 50, 55, 59, 64}}

func CreateInstrument(t InstrumentType) Instrument {
	//for now, everyone gets standard tuning
	return &guitarStandardTuning;
}