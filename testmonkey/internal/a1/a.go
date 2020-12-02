package a1

type MyStruct struct {
	a string
	b string
}

func (t *MyStruct) hello() string {
	return t.a+t.b
}

func New(a, b string) *MyStruct {
	return &MyStruct{
		a: a,
		b: b,
	}
}
