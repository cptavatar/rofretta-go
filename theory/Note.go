package theory

import (
	"bytes"
	"strings"
)

// Notes capture the natural name of a note (A,B,C,D,E,F,G) as well
// as how sharp/flat it is. We can also send a "Disabled" note as a placeholder
// if, for instance, we are sending back the note names of the current voicing
// but not all strings are played
type Note interface {
	Natural() NoteName
	Modifier() int8
	Disabled() bool
	Name() string
	EnharmonicValue() int8
}

type note struct {
	natural  NoteName
	modifier int8
}

func (note note) Natural() NoteName {
	return note.natural
}
func (note note) Modifier() int8 {
	return note.modifier
}

func (note note) Disabled() bool {
	return note.natural.Offset() == Disabled
}

func (note note) EnharmonicValue() int8 {
	if (note.Disabled()) {
		return -1
	}
	var calculatedValue int8 = note.natural.Offset() + note.modifier

	if calculatedValue < 0 {
		calculatedValue = 12 + calculatedValue
	} else if calculatedValue > 11 {
		calculatedValue = calculatedValue - 12
	}

	return calculatedValue
}

func (note note) Name() string {
	var buffer bytes.Buffer
	buffer.WriteString(note.natural.Name())
	for i := note.modifier; i > 0; i-- {
		buffer.WriteRune('\u266F')
	}
	for i := note.modifier; i < 0; i++ {
		buffer.WriteRune('\u266D')
	}
	return buffer.String()
}

func CreateNote(str string) Note {
	var nt note
	nt.natural = CreateNoteName("")
	tempStr := []rune(strings.TrimSpace(str))
	for i := 0; i < len(tempStr); i++ {
		switch (tempStr[i]) {
		case 'C':
			fallthrough
		case 'D':
			fallthrough
		case 'E':
			fallthrough
		case 'F':
			fallthrough
		case 'G':
			fallthrough
		case 'A':
			fallthrough
		case 'B':
			nt.natural = CreateNoteName(string(tempStr[i]))
			break;
		case '\u266F':
			fallthrough
		case '#':
			nt.modifier++;
			break;
		case '\u266D':
			fallthrough
		case 'b':
			nt.modifier--;
			break;
		}
	}
	return nt;
}

func CreateNoteInt(midiValue int8) Note {
	baseValue := midiValue % 12;
	var note note
	note.natural = CreateNoteNameByOffset(baseValue);

	if note.natural.Index() == -1 {
		note.modifier = 1
		note.natural = CreateNoteNameByOffset(baseValue - 1);
	}
	return note
}

func CreateNoteForInterval(interval Interval, root Note) Note {
	index := (root.Natural().Index() + interval.DiatonicSteps()) % 7
	noteName := CreateNoteNameByIndex(index)
	newNote := note{noteName, 0}
	baseValue := newNote.EnharmonicValue()
	if (baseValue < root.EnharmonicValue()) {
		baseValue = baseValue + 12
	}
	newNote.modifier = interval.Offset() - (baseValue - root.EnharmonicValue())
	return newNote
}
