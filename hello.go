package main

import (
	"fmt"
	"stringutil"
)

func main() {
	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

}
