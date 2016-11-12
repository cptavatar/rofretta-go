package theory

import "strings"

type NoteName interface {
	Offset() int8
	Name() string
	Index() int8
}

type noteName struct {
	offset int8
	name   string
	index  int8
}

var notesByName = make(map[string]noteName)
var notesByIndex = make(map[int8]noteName)
var notesByOffset = make(map[int8]noteName)

func init() {
	var names = []noteName{
		noteName{name:"", offset: -1, index:-1},
		noteName{name:"C", offset: 0, index:0},
		noteName{name:"D", offset: 2, index:1},
		noteName{name:"E", offset: 4, index:2},
		noteName{name:"F", offset: 5, index:3},
		noteName{name:"G", offset: 7, index:4},
		noteName{name:"A", offset: 9, index:5},
		noteName{name:"B", offset: 11, index:6}}
	for i := 0; i < 8; i++ {
		notesByOffset[names[i].offset] = names[i]
		notesByIndex[names[i].index] = names[i]
		notesByName[names[i].name] = names[i]
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

func CreateNoteByName(name string) NoteName {
	if val, ok := notesByName[strings.ToUpper(name)]; ok {
		return val
	}
	return notesByName[""]
}
func CreateNoteByIndex(idx int8) NoteName {
	if val, ok := notesByIndex[idx]; ok {
		return val
	}
	return notesByIndex[-1]
}
func CreateNoteByOffset(offset int8) NoteName {
	if val, ok := notesByOffset[offset]; ok {
		return val

	}
	return notesByOffset[-1]
}