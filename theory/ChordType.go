package theory

type ChordType uint8

const (
	Major ChordType = iota
	Minor
	DominantSeven
	Suspended
	MinorSeven
	MajorFlatFive
	Augmented
	MajorSix
	MajorSixAddNine
	MinorSix
	MinorSixAddNine
	SevenSuspended
	MinorSevenFlatFive
	DiminishedSeven
	SevenSharpFive
	SevenFlatFive
	MajorSeven
	MinorMajorSeven
	MajorAddNine
	SevenSharpNine
	SevenFlatNine
	SevenSharpFlatNine
	MinorNine
	NineSharpFive
	NineFlatFive
	MajorNine
	NineSharpEleven
	MinorNineMajorSeven
	Eleventh
	MinorEleven
	Thirteenth
	ThirteenthFlatNine
	ThirteenthFlatFiveFlatNine
	MinorThirteen
)

type Chord interface {
	Name() string
	ShortName() string
	Intervals() []Interval
}

type chord struct {
	name      string
	shortName string
	intervals []IntervalType
}

var
	MAJOR("Major", "maj", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH)),
MINOR("Minor", "min", Arrays.asList(MINOR_THIRD, PERFECT_FIFTH)),
DOMINANT_SEVEN("Dominant Seven", "7", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MINOR_SEVENTH)),
SUSPENDED("Suspended", "sus", Arrays.asList(PERFECT_FOURTH, PERFECT_FIFTH)),
MINOR_SEVEN("Minor Seven", "min7", Arrays.asList(MINOR_THIRD, PERFECT_FIFTH, MINOR_SEVENTH)),

MAJOR_FLAT_FIVE("Major Flat Five", "maj-5", Arrays.asList(MAJOR_THIRD, TRITONE)),
AUGMENTED("Augmented", "+", Arrays.asList(MAJOR_THIRD, AUGMENTED_FIFTH)),
MAJOR_SIX("Major Six", "6", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SIXTH)),
MAJOR_SIX_ADD_NINE("Major Six Add Nine", "6/9", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SIXTH, MAJOR_NINTH)),
MINOR_SIX("Minor Six", "m6", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SIXTH)),
MINOR_SIX_ADD_NINE("Minor Six Add Nine", "m6/9", Arrays.asList(MINOR_THIRD, PERFECT_FIFTH, MAJOR_SIXTH, MAJOR_NINTH)),


SEVEN_SUSPENDED("Seven Suspended", "sus7", Arrays.asList(PERFECT_FOURTH, PERFECT_FIFTH, MINOR_SEVENTH)),

MINOR_SEVEN_FLAT_FIVE("Minor Seven Flat Five", "min7-5", Arrays.asList(MINOR_THIRD, TRITONE, MINOR_SEVENTH)),
DIMINISHED_SEVEN("Diminished Seven", "dim7", Arrays.asList(MINOR_THIRD, TRITONE, DIMINISHED_SEVENTH)),
SEVEN_SHARP_FIVE("Seven Sharp Five", "7+5", Arrays.asList(MAJOR_THIRD, AUGMENTED_FIFTH, MINOR_SEVENTH)),
SEVEN_FLAT_FIVE("Seven Flat Five", "7-5", Arrays.asList(MAJOR_THIRD, TRITONE, MINOR_SEVENTH)),
MAJOR_SEVEN("Major Seven", "maj7", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH)),
MINOR_MAJOR_SEVEN("Minor Major Seven", "min+7", Arrays.asList(MINOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH)),
MAJOR_ADD_NINE("Major Add Nine", "add9", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_NINTH)),
SEVEN_SHARP_NINE("Seven Sharp Nine", "7+9", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MINOR_SEVENTH, AUGMENTED_NINTH)),
SEVEN_FLAT_NINE("Seven Flat Nine", "7-9", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MINOR_SEVENTH, MINOR_NINTH)),
SEVEN_SHARP_FIVE_FLAT_NINE("Seven Sharp Five Flat Nine", "7+5-9", Arrays.asList(MAJOR_THIRD, AUGMENTED_FIFTH, MINOR_SEVENTH, MINOR_NINTH)),
NINTH("Ninth", "9", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MINOR_SEVENTH, MAJOR_NINTH)),
MINOR_NINE("Minor Nine", "min9", Arrays.asList(MINOR_THIRD, PERFECT_FIFTH, MINOR_SEVENTH, MAJOR_NINTH)),
NINE_SHARP_FIVE("Nine Sharp Five", "9+5", Arrays.asList(MAJOR_THIRD, AUGMENTED_FIFTH, MINOR_SEVENTH, MAJOR_NINTH)),
NINE_FLAT_FIVE("Nine Flat Five", "9-5", Arrays.asList(MAJOR_THIRD, TRITONE, MINOR_SEVENTH, MAJOR_NINTH)),
MAJOR_NINE("Major Nine", "maj9", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH)),
NINE_SHARP_ELEVEN("Nine Sharp Eleven", "9+11", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH)),
MINOR_NINE_MAJOR_SEVEN("Minor Nine Major Seven", "min9+7", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH)),
ELEVENTH("Eleventh", "11", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH)),
MINOR_ELEVEN("Minor Eleven", "min11", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH)),
THIRTEENTH("Thirteenth", "13", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH)),
THIRTEENTH_FLAT_NINE("Thirteenth Flat Nine", "13-9", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH)),
THIRTEENTH_FLAT_FIVE_FLAT_NINE("Thirteenth Flat Five Flat Nine", "13-5-9", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH)),
MINOR_THIRTEENTH("Minor Thirteenth", "min13", Arrays.asList(MAJOR_THIRD, PERFECT_FIFTH, MAJOR_SEVENTH, MAJOR_NINTH));

private static final Map<String, ChordType> typeByPrettyName = new HashMap<>();
private static final Map<String, ChordType> typeByAbbreviation = new HashMap<>();
private static final Map<Integer, List<ChordType>> typesBySize = new HashMap<>();

static {
for (ChordType type : ChordType.values()) {
typeByPrettyName.put(type.getPrettyName(), type);
typeByAbbreviation.put(type.getAbbreviation(), type);
int size = type.intervals.size() + 1;
if (typesBySize.containsKey(size)) {
typesBySize.get(size).add(type);
} else {
List<ChordType> types = new ArrayList<>();
types.add(type);
typesBySize.put(size, types);
}
}
}


private final String prettyName;
private final String abbreviation;
private final List<Interval> intervals;
}
