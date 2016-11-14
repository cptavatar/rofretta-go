package theory

type IntervalType uint8

const (
	ITPerfectUnision IntervalType = iota
	ITMinorSecond
	ITMajorSecond
	ITMinorThird
	ITMajorThird
	ITPerfectFourth
	ITTritone
	ITPerfectFifth
	ITAugmentedFifth
	ITMinorSixth
	ITMajorSixth
	ITDiminishedSeventh
	ITMinorSeventh
	ITMajorSeventh
	ITOctive
	ITMinorNinth
	ITMajorNinth
	ITAugmentedNinth
	ITMinorTenth
	ITMajorTenth
	ITPerfectEleventh
	ITAugmentedEleventh
	ITPerfectTwelfth
	ITMinorThirteen
	ITMajorThirteen
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

var intervalsByType = make(map[IntervalType]*interval)

func init() {
	intervalsByType[ITPerfectUnision] = &interval{"Unison", 0, "0", 0}
	intervalsByType[ITMinorSecond] = &interval{"Minor Second", 1, "\u266D2", 1}
	intervalsByType[ITMajorSecond] = &interval{"Major second", 2, "2", 1}
	intervalsByType[ITMinorThird] = &interval{"Minor Third", 3, "\u266D3", 2}
	intervalsByType[ITMajorThird] = &interval{"Major Third", 4, "3", 2}
	intervalsByType[ITPerfectFourth] = &interval{"Perfect Fourth", 5, "4", 3}
	intervalsByType[ITTritone] = &interval{"Tritone", 6, "\u266D5", 4}
	intervalsByType[ITPerfectFifth] = &interval{"Perfect Fifth", 7, "5", 4}
	intervalsByType[ITAugmentedFifth] = &interval{"Augmented Fifth", 8, "\u266F5", 4}
	intervalsByType[ITMinorSixth] = &interval{"Minor Sixth", 8, "\u266D6", 5}
	intervalsByType[ITMajorSixth] = &interval{"Major Sixth", 9, "6", 5}
	intervalsByType[ITDiminishedSeventh] = &interval{"Diminished Seventh", 9, "\u266D\u266D7", 6}
	intervalsByType[ITMinorSeventh] = &interval{"Minor Seventh", 10, "\u266D7", 6}
	intervalsByType[ITMajorSeventh] = &interval{"Major Seventh", 11, "7", 6}
	intervalsByType[ITOctive] = &interval{"Octive", 12, "8", 7}
	intervalsByType[ITMinorNinth] = &interval{"Minor Ninth", 13, "\u266D9", 8}
	intervalsByType[ITMajorNinth] = &interval{"Major Ninth", 14, "9", 8}
	intervalsByType[ITAugmentedNinth] = &interval{"Augmented Ninth", 15, "\u266F9", 8}
	intervalsByType[ITMinorTenth] = &interval{"Minor Tenth", 15, "\u266D10", 9}
	intervalsByType[ITMajorTenth] = &interval{"Major Tenth", 16, "10", 9}
	intervalsByType[ITPerfectEleventh] = &interval{"Perfect Eleventh", 17, "11", 10}
	intervalsByType[ITAugmentedEleventh] = &interval{"Augmented Eleventh", 18, "\u266F11", 10}
	intervalsByType[ITPerfectTwelfth] = &interval{"Perfect Twelfth", 19, "12", 11}
	intervalsByType[ITMinorThirteen] = &interval{"Minor Thirteenth", 20, "\u266D13", 12}
	intervalsByType[ITMajorThirteen] = &interval{"Major Thirteenth", 21, "13", 12}
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
