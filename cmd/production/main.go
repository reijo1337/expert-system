package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	backward = flag.Bool("backward", false, "Прямой ход")
	fileName = flag.String("f", "sas.json", "Файл с правилами")
)

func main() {
	flag.Parse()
	fmt.Println(*backward)
	fmt.Println(os.Args)
	flag.PrintDefaults()
}
