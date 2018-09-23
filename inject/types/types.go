package types

type t struct {
	value string
}

func (t t) String() string { return t.value }

var (
	Templates      = t{"templates"}
	FileServer     = t{"fileserver"}
	PostRepository = t{"postrepository"}
	RootController = t{"rootcontroller"}
	PostController = t{"postcontroller"}
)
