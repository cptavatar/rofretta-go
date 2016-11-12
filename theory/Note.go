package theory

type Note interface {
	Root() NoteName
	Modifer() int8
}

type note struct {
	root     NoteName
	modifier int8
}

