package main

import (
	"fmt"

	"github.com/alexamantea/go/mathutil"
	"github.com/alexamantea/go/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.Reverse("!iG illeH"))
	fmt.Println(stringutil.Reverse("!erdnaxelA"))

	fmt.Println(stringutil.WordCount("i wonder if this is correct is it"))
	fmt.Println(stringutil.WordCount("i am not sure if this is working is it well working"))

	fmt.Println(mathutil.Sqrt(244))
	fmt.Println(mathutil.Sqrt(4))
	fmt.Println(mathutil.Sqrt(100))
}
