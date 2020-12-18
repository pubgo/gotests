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
	HH string
}

func (t hh) init1() string {
	return t.HH
}

func Newhh(s string) *hh {
	return &hh{HH: s}
}
