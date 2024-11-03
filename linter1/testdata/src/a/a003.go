package a

import "fmt"

func f4() {
	b1 := true
	var b2 = false

	fmt.Println(b1, b2)
}

func f5(b3, b31 bool) (bool, error) {}
func f6() (b4 bool, err error)      {}

type int1 interface {
	f11(b6 bool)
	f12() (b7 bool)
}

type str1 struct {
	b5 bool
}

func (t *str1) hello(b8 bool)    {}
func (t *str1) world() (b9 bool) {}

func f7() {
	f8 := func(b10 bool) {}
	f9 := func() (b11 bool) {}
	var f8 = func(b12 bool) {}
	var f9 = func() (b13 bool) {}
}
