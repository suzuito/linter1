package a

import "fmt"

var v001 = true    // want "a boolean variable does not match pattern"
const v002 = false // want "a boolean variable does not match pattern"

var v012, uuu = true, "this is a test" // want "a boolean variable does not match pattern"

var foo = "this is a test"

var isHoge = true

const hasHoge = true

type s001 struct {
	v003    bool // want "a boolean variable does not match pattern"
	areHoge bool
}

func f(
	v004 bool, // want "a boolean variable does not match pattern"
	isHoge bool,
) (
	v005 bool, // want "a boolean variable does not match pattern"
	hasHoge bool,
	err error,
) {
	v006 := false // want "a boolean variable does not match pattern"
	areHoge := true

	v007 := v006 && areHoge // want "a boolean variable does not match pattern"
	hasFoo := v006 && areHoge
	fmt.Println(v007, hasFoo)
	return v006, areHoge, nil
}
