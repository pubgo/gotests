package main

func main() {
	init1(func(i int) {

	})
}

func init1(fn func(i int)) {
	//val=yield fn(init2())
}

func init2() int {
	return 0
}
