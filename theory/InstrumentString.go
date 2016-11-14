package theory

type InstrumentString interface {
	Fret() int8
	Finger() Finger
	Disabled() bool
}

type instrumentString struct {
	fret   int8
	finger Finger
}

func CreateInstrumentString(fret int8, finger Finger) InstrumentString {
	return instrumentString{
		fret: fret,
		finger: finger}
}

func CreateInstrumentStringInt(fret int8, finger int) InstrumentString {
	return CreateInstrumentString(fret, CreateFinger(finger))
}

func (string instrumentString) Fret() int8 {
	return string.fret
}

func (string instrumentString) Finger() Finger {
	return string.finger
}

func (string instrumentString) Disabled() bool {
	return string.fret == -1
}