package theory

import "strings"

// The natural notes (white keys). The offset can be used to figure actual space beteween
// notes (taking in account the black keys), while index can be used for diatonic calculations.
type NoteName interface {
	Name() string
	Offset() int8
	Index() int8
}

type noteName struct {
	name   string
	offset int8
	index  int8
}

var notesByName = make(map[string]*noteName)
var notesByIndex = make(map[int8]*noteName)
var notesByOffset = make(map[int8]*noteName)

func init() {
	var names = []noteName{
		noteName{"", -1, -1},
		noteName{"C", 0, 0},
		noteName{"D", 2, 1},
		noteName{"E", 4, 2},
		noteName{"F", 5, 3},
		noteName{"G", 7, 4},
		noteName{"A", 9, 5},
		noteName{"B", 11, 6}}
	for i := 0; i < 8; i++ {
		notesByOffset[names[i].offset] = &names[i]
		notesByIndex[names[i].index] = &names[i]
		notesByName[names[i].name] = &names[i]
	}

}
func (nn noteName) Offset() int8 {
	return nn.offset
}
func (nn noteName) Name() string {
	return nn.name
}
func (nn noteName) Index() int8 {
	return nn.index
}

func CreateNoteName(name string) NoteName {
	if val, ok := notesByName[strings.ToUpper(name)]; ok {
		return val
	}
	return notesByName[""]
}
func CreateNoteNameByIndex(idx int8) NoteName {
	if val, ok := notesByIndex[idx]; ok {
		return val
	}
	return notesByIndex[-1]
}
func CreateNoteNameByOffset(offset int8) NoteName {
	if val, ok := notesByOffset[offset]; ok {
		return val
	}
	return notesByOffset[-1]
}