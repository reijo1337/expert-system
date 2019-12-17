package main

import "fmt"

func main() {
	fmt.Println(getTokenList(
		`brother_sister(joe, monica).
		brother_sister(eric, erica).
		brother_sister(jim, rebecca).`,
	))
}
