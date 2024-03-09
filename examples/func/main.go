package main

import (
	"fmt"

	cs "github.com/ortense/consolestyle"
)

func main() {
	fmt.Println(cs.Italic(cs.Green("\"Simplicity is the ultimate sophistication.\"")))
	fmt.Println(cs.Dim("- Leonardo da Vinci"))
}
