package a

import "fmt"

func f1() {
	_ = append([]int{1, 2, 3})
}

func f2() {
	return
	f1 := 1
	{
		f1 := "hoge"
		fmt.Println(f1)
	}
	fmt.Println(f1)
}
