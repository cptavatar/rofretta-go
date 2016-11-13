package theory

import "strings"

type Note interface {
	Root() NoteName
	Modifier() int8
	Disabled() bool
}

type note struct {
	root     NoteName
	modifier int8
}

func (note note) Root() NoteName {
	return note.root
}
func (note note) Modifier() int8 {
	return note.modifier
}
func (note note) Disabled() bool {
	return note.root.Offset() == -1
}

func CreateNote(str string) Note {
	var nt note
	nt.root = CreateNoteName("")
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
			nt.root = CreateNoteName(string(tempStr[i]))
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
	note.root = CreateNoteNameByOffset(baseValue);

	if note.root.Index() == -1 {
		note.modifier = 1
		note.root = CreateNoteNameByOffset(baseValue - 1);
	}
	return note
}
/*
public static List<Note> notes(ChordType type, Note root) {
List<Interval> intervals = type.getIntervals();
List<Note> notes = new ArrayList<>(intervals.size() + 1);
notes.add(root);
intervals.forEach(interval -> notes.add(note(interval, root)));
return notes;
}

public static Note note(Interval interval, Note root) {
NoteName name = root.getRoot();
int index = (name.getIndex() + interval.getDiatonicSteps()) % 7;
NoteName newName = NoteName.valueOfIndex(index);
int rootValue = NoteService.enharmonicValue(root);
int newNoteBaseValue = NoteService.enharmonicValue(new Note(newName, 0));
if (newNoteBaseValue < rootValue) {
newNoteBaseValue = newNoteBaseValue + 12;
}
int modifier = interval.offset() - (newNoteBaseValue - rootValue);
return new Note(newName, modifier);

}*/