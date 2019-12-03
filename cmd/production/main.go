package main

import (
	"flag"
	"fmt"
)

var (
	backward = flag.Bool("backward", false, "Прямой ход")
	fileName = flag.String("f", "sas.json", "Файл с правилами")
)

func main() {
	rules, err := loadRules(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("rules:")
	for k, v := range rules {
		fmt.Printf("%s: %v\n", k, v)
	}
}
