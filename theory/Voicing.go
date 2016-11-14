package theory

type UnidentifedVoicing interface {
	Strings() []InstrumentString
	Instrument() Instrument
}

type Voicing interface {
	UnidentifedVoicing
	Type() VocingType
	Root() Note
	Chord() Chord
}

type unidentifedVoicing struct {
	strings    []InstrumentString
	instrument Instrument
}

type voicing struct {
	unidentifedVoicing
	vType VocingType
	root  Note
	chord Chord
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
func (vc voicing) Chord() Chord {
	return vc.chord
}

func CreateVoicing(strings []InstrumentString, instrument Instrument, root Note, chord Chord) Voicing {
	retval := voicing{root:root, chord:chord }
	retval.strings = strings
	retval.instrument = instrument

	//This algorithm could use some work, only really works for open
	var muted, open bool
	for _, str := range retval.Strings() {
		if str.Fret() == 0 {
			open = true
		} else if str.Fret() == -1 {
			muted = true
		}
	}
	if open {
		retval.vType = Open
	} else if muted {
		retval.vType = Shape
	} else {
		retval.vType = Bar
	}

	return retval
}
func CreateUnidentifedVoicing(strings []InstrumentString, instrument Instrument) UnidentifedVoicing {
	return unidentifedVoicing{strings, instrument}
}