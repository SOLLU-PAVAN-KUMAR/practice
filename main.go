package main

import (
	"fmt"
	ut "go/project/stringutil"
)

func main() {
	s := "pavan"
	s = ut.Upper(s)
	fmt.Println(s)
	fmt.Println(ut.Lower(s))
}
