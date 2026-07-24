package creatures

type Animal struct {
	Name, Class, Emoji string
	avgLifespan        int // example of a private struct field: begins with a lowercase letter
	Domestic           bool
	Human              human
}

type human struct {
	name string
	age  int
}
