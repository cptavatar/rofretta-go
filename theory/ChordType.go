package theory

// ChordType are constants that can be used to indicate the kind of chord one wants
type ChordType uint8

const (
	CTMajor ChordType = iota
	CTMinor
	CTDominantSeven
	CTSuspended
	CTMinorSeven
	CTMajorFlatFive
	CTAugmented
	CTMajorSix
	CTMajorSixAddNine
	CTMinorSix
	CTMinorSixAddNine
	CTSevenSuspended
	CTMinorSevenFlatFive
	CTDiminishedSeven
	CTSevenSharpFive
	CTSevenFlatFive
	CTMajorSeven
	CTMinorMajorSeven
	CTMajorAddNine
	CTSevenSharpNine
	CTSevenFlatNine
	CTSevenSharpFlatNine
	CTNinth
	CTMinorNine
	CTNineSharpFive
	CTNineFlatFive
	CTMajorNine
	CTNineSharpEleven
	CTMinorNineMajorSeven
	CTEleventh
	CTMinorEleven
	CTThirteenth
	CTThirteenthFlatNine
	CTThirteenthFlatFiveFlatNine
	CTMinorThirteen
)

//A chord is made up of a long descriptive name, a short
//name as we would expect in a file, and a list of intervals that define the chord
type Chord interface {
	Name() string
	ShortName() string
	Intervals() []Interval
	Notes(root Note) []Note
}

type chord struct {
	name      string
	shortName string
	intervals []IntervalType
}

func (chord chord) Name() string {
	return chord.name
}
func (chord chord) ShortName() string {
	return chord.shortName
}
func (chord chord) Intervals() []Interval {
	intervalSlice := make([]Interval, len(chord.intervals) + 1, len(chord.intervals) + 1)
	intervalSlice[0] = CreateInterval(ITPerfectUnision)
	for i := 1; i < len(intervalSlice); i++ {
		intervalSlice[i] = CreateInterval(chord.intervals[i - 1])
	}
	return intervalSlice
}
func (chord chord) Notes(root Note) []Note {
	intervals := chord.Intervals()
	size := len(intervals)
	notes := make([]Note, size, size)
	for i, interval := range intervals {
		notes[i] = CreateNoteForInterval(interval, root)
	}
	return notes
}

var chordByType = make(map[ChordType]*chord)
var chordByShortName = make(map[string]*chord)
var chordsByIntervalCount = make(map[int][]Chord)

//TODO chordtype should be part of the struct, testable, returnable
func init() {
	chordByType[CTMajor] = &chord{"Major", "maj", []IntervalType{ITMajorThird, ITPerfectFifth}}
	chordByType[CTMinor] = &chord{"Minor", "min", []IntervalType{ITMinorThird, ITPerfectFifth}}
	chordByType[CTDominantSeven] = &chord{"Dominant Seven", "7", []IntervalType{ITMajorThird, ITPerfectFifth, ITPerfectFifth}}
	chordByType[CTSuspended] = &chord{"Suspended", "sus", []IntervalType{ITPerfectFourth, ITPerfectFifth}}
	chordByType[CTMinorSeven] = &chord{"Minor Seven", "min7", []IntervalType{ITMinorThird, ITPerfectFifth, ITMinorSeventh}}
	chordByType[CTMajorFlatFive] = &chord{"Major Flat Five", "maj-5", []IntervalType{ITMajorThird, ITTritone}}
	chordByType[CTAugmented] = &chord{"Augmented", "+", []IntervalType{ITMajorThird, ITAugmentedFifth}}
	chordByType[CTMajorSix] = &chord{"Major Six", "6", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSixth}}
	chordByType[CTMajorSixAddNine] = &chord{"Major Six Add Nine", "6/9", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSixth, ITMajorNinth}}
	chordByType[CTMinorSix] = &chord{"Minor Six", "m6", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSixth}}
	chordByType[CTMinorSixAddNine] = &chord{"Minor Six Add Nine", "m6/9", []IntervalType{ITMinorThird, ITPerfectFifth, ITMajorSixth, ITMajorNinth}}
	chordByType[CTSevenSuspended] = &chord{"Seven Suspended", "sus7", []IntervalType{ITPerfectFourth, ITPerfectFifth, ITMinorSeventh}}
	chordByType[CTMinorSevenFlatFive] = &chord{"Minor Seven Flat Five", "min7-5", []IntervalType{ITMinorThird, ITTritone, ITMinorSeventh}}
	chordByType[CTDiminishedSeven] = &chord{"Diminished Seven", "dim7", []IntervalType{ITMinorThird, ITTritone, ITDiminishedSeventh}}
	chordByType[CTSevenSharpFive] = &chord{"Seven Sharp Five", "7+5", []IntervalType{ITMajorThird, ITAugmentedFifth, ITMinorSeventh}}
	chordByType[CTSevenFlatFive] = &chord{"Seven Flat Five", "7-5", []IntervalType{ITMajorThird, ITTritone, ITMinorSeventh}}
	chordByType[CTMajorSeven] = &chord{"Major Seven", "maj7", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh}}
	chordByType[CTMinorMajorSeven] = &chord{"Minor Major Seven", "min+7", []IntervalType{ITMinorThird, ITPerfectFifth, ITMajorSeventh}}
	chordByType[CTMajorAddNine] = &chord{"Major Add Nine", "add9", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorNinth}}
	chordByType[CTSevenSharpNine] = &chord{"Seven Sharp Nine", "7+9", []IntervalType{ITMajorThird, ITPerfectFifth, ITMinorSeventh, ITAugmentedNinth}}
	chordByType[CTSevenFlatNine] = &chord{"Seven Flat Nine", "7-9", []IntervalType{ITMajorThird, ITPerfectFifth, ITMinorSeventh, ITMinorNinth}}
	chordByType[CTSevenSharpFlatNine] = &chord{"Seven Sharp Five Flat Nine", "7+5-9", []IntervalType{ITMajorThird, ITAugmentedFifth, ITMinorSeventh, ITMinorNinth}}
	chordByType[CTNinth] = &chord{"Ninth", "9", []IntervalType{ITMajorThird, ITPerfectFifth, ITMinorSeventh, ITMajorNinth}}
	chordByType[CTMinorNine] = &chord{"Minor Nine", "min9", []IntervalType{ITMinorThird, ITPerfectFifth, ITMinorSeventh, ITMajorNinth}}
	chordByType[CTNineSharpFive] = &chord{"Nine Sharp Five", "9+5", []IntervalType{ITMajorThird, ITAugmentedFifth, ITMinorSeventh, ITMajorNinth}}
	chordByType[CTNineFlatFive] = &chord{"Nine Flat Five", "9-5", []IntervalType{ITMajorThird, ITTritone, ITMinorSeventh, ITMajorNinth}}
	chordByType[CTMajorNine] = &chord{"Major Nine", "maj9", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}
	chordByType[CTNineSharpEleven] = &chord{"Nine Sharp Eleven", "9+11", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}
	chordByType[CTMinorNineMajorSeven] = &chord{"Minor Nine Major Seven", "min9+7", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}
	chordByType[CTEleventh] = &chord{"Eleventh", "11", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}
	chordByType[CTMinorEleven] = &chord{"Minor Eleven", "min11", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}
	chordByType[CTThirteenth] = &chord{"Thirteenth", "13", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}
	chordByType[CTThirteenthFlatNine] = &chord{"Thirteenth Flat Nine", "13-9", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}
	chordByType[CTThirteenthFlatFiveFlatNine] = &chord{"Thirteenth Flat Five Flat Nine", "13-5-9", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}
	chordByType[CTMinorThirteen] = &chord{"Minor Thirteenth", "min13", []IntervalType{ITMajorThird, ITPerfectFifth, ITMajorSeventh, ITMajorNinth}}

	for _, chrd := range chordByType {
		chordByShortName[chrd.shortName] = chrd
		intervalCount := len(chrd.intervals) + 1
		if slice, ok := chordsByIntervalCount[intervalCount]; ok {
			chordsByIntervalCount[intervalCount] = append(slice, chrd)
		} else {
			chrdSlice := make([]Chord, 1, 7)
			chrdSlice[0] = chrd
			chordsByIntervalCount[intervalCount] = chrdSlice
		}
	}
}
// Given a short name, the chord for the name.
// TODO - consider better error handling, bad names are currently fatal
func CreateChordByShortName(shortName string) Chord {
	retval := chordByShortName[shortName]
	if (retval == nil) {
		Log.WithField("name", shortName).Fatal("Unknown chord type")
	}
	return retval
}

func CreateChordsByIntervalCount(count int) []Chord {
	return chordsByIntervalCount[count]
}
// Given a ChordType, return a chord definition
func CreateChord(tp ChordType) Chord {
	return chordByType[tp]
}

