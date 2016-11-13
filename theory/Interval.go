package theory

type IntervalType uint8

const (
	PerfectUnision IntervalType = iota
	MinorSecond
	MajorSecond
	MinorThird
	MajorThird
	PerfectFourth
	Tritone
	PerfectFifth
	AugmentedFifth
	MinorSixth
	MajorSixth
	DiminishedSeventh
	MinorSeventh
	MajorSeventh
	Octive
	MinorNinth
	MajorNinth
	AugmentedNinth
	MinorTenth
	MajorTenth
	PerfectEleventh
	AugmentedEleventh
	PerfectTwelfth
	MinorThirteen
	MajorThirteen
)

type Interval interface {
	Name() string
	Offset() uint8
	ShortName() string
	DiatonicSteps() uint8
}

type interval struct {
	name          string
	offset        uint8
	shortName     string
	diatonicSteps uint8
}

var intervalsByType = make(map[IntervalType]interval)

func init() {
	intervalsByType[PerfectUnision] = interval{"Unison", 0, "0", 0}
	intervalsByType[MinorSecond] = interval{"Minor Second", 1, "\u266D2", 1}
	intervalsByType[MajorSecond] = interval{"Major second", 2, "2", 1}
	intervalsByType[MinorThird] = interval{"Minor Third", 3, "\u266D3", 2}
	intervalsByType[MajorThird] = interval{"Major Third", 4, "3", 2}
	intervalsByType[PerfectFourth] = interval{"Perfect Fourth", 5, "4", 3}
	intervalsByType[Tritone] = interval{"Tritone", 6, "\u266D5", 4}
	intervalsByType[PerfectFifth] = interval{"Perfect Fifth", 7, "5", 4}
	intervalsByType[AugmentedFifth] = interval{"Augmented Fifth", 8, "\u266F5", 4}
	intervalsByType[MinorSixth] = interval{"Minor Sixth", 8, "\u266D6", 5}
	intervalsByType[MajorSixth] = interval{"Major Sixth", 9, "6", 5}
	intervalsByType[DiminishedSeventh] = interval{"Diminished Seventh", 9, "\u266D\u266D7", 6}
	intervalsByType[MinorSeventh] = interval{"Minor Seventh", 10, "\u266D7", 6}
	intervalsByType[MajorSeventh] = interval{"Major Seventh", 11, "7", 6}
	intervalsByType[Octive] = interval{"Octive", 12, "8", 7}
	intervalsByType[MinorNinth] = interval{"Minor Ninth", 13, "\u266D9", 8}
	intervalsByType[MajorNinth] = interval{"Major Ninth", 14, "9", 8}
	intervalsByType[AugmentedNinth] = interval{"Augmented Ninth", 15, "\u266F9", 8}
	intervalsByType[MinorTenth] = interval{"Minor Tenth", 15, "\u266D10", 9}
	intervalsByType[MajorTenth] = interval{"Major Tenth", 16, "10", 9}
	intervalsByType[PerfectEleventh] = interval{"Perfect Eleventh", 17, "11", 10}
	intervalsByType[AugmentedEleventh] = interval{"Augmented Eleventh", 18, "\u266F11", 10}
	intervalsByType[PerfectTwelfth] = interval{"Perfect Twelfth", 19, "12", 11}
	intervalsByType[MinorThirteen] = interval{"Minor Thirteenth", 20, "\u266D13", 12}
	intervalsByType[MajorThirteen] = interval{"Major Thirteenth", 21, "13", 12}
}

func (interval interval) Name() string {
	return interval.name
}
func (interval interval) Offset() uint8 {
	return interval.offset
}
func (interval interval) ShortName() string {
	return interval.shortName
}
func (interval interval) DiatonicSteps() uint8 {
	return interval.diatonicSteps
}

func CreateInterval(it IntervalType) Interval {
	return intervalsByType[it]
}
