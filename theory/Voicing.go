package theory

type UnidentifedVoicing interface {
	Strings() []InstrumentString
	Instrument() Instrument
}

type Voicing interface {
	UnidentifedVoicing
	Type() VocingType
	Root() Note
	Chord() ChordType
}

type unidentifedVoicing struct {
	strings    []InstrumentString
	instrument Instrument
}

type voicing struct {
	unidentifedVoicing
	vType VocingType
	root  Note
	chord ChordType
}

func (vc unidentifedVoicing) Strings() []InstrumentString {
	return vc.strings
}
func (vc unidentifedVoicing) Instrument() Instrument {
	return vc.instrument
}
func (vc voicing) Type() VocingType {
	return vc.vType
}
func (vc voicing) Root() Note {
	return vc.root
}
func (vc voicing) Chord() ChordType {
	return vc.chord
}
