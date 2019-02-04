package main

import (
	"fmt"

	"github.com/alexamantea/go/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.Reverse("!iG illeH"))
	fmt.Println(stringutil.Reverse("!erdnaxelA"))

	fmt.Println(stringutil.WordCount("sera que acertei que nao acertei"))
	fmt.Println(stringutil.WordCount("eu nao sei quantas palavras tem aqui"))
}
