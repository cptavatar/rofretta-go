package theory

// Finger is used to indicate which finger is used to press a fret, if we know it
type Finger int8

const (
	Unknown Finger = -1
	Index Finger = 1
	Middle Finger = 2
	Ring Finger = 3
	Little Finger = 4
	Thumb Finger = 5
)

// Given an int, return the appropriate Finger if known else Unknown
func CreateFinger(i int) Finger {
	switch (i){
	case 1:
		return Index
	case 2:
		return Middle
	case 3:
		return Ring
	case 4:
		return Little
	case 5:
		return Thumb
	default:
		return Unknown
	}
}
