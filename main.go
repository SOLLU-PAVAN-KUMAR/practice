package main

import (
	"fmt"
	ut "github.com/SOLLU-PAVAN-KUMAR/stringutil"
)

func main() {
	s := "pavan"
	s = ut.Upper(s)
	fmt.Println(s)
	fmt.Println(ut.Lower(s))
}
