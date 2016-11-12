package theory

type Finger int8

const (
	Unknown Finger = -1
	Index Finger = 1
	Middle Finger = 2
	Ring Finger = 3
	Little Finger = 4
	Thumb Finger = 5
)

func FingerById(i int) Finger {
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
