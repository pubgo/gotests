package a1

type MyStruct struct {
	a string
	b string
}

func (t *MyStruct) hello() string {
	return t.a + t.b
}

func (t *MyStruct) Hello() string {
	return t.hello()
}

func New(a, b string) *MyStruct {
	return &MyStruct{
		a: a,
		b: b,
	}
}

type hh struct {
	h1 string
}

func (t hh) init1() string {
	return t.h1
}

func Newhh(s string) *hh {
	return &hh{h1: s}
}
