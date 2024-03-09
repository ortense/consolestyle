package main

import (
	"fmt"

	"github.com/ortense/consolestyle"
)

func main() {
	message := consolestyle.
		Style("Hey there! 👋").Cyan().Italic().
		EmptyLine().
		NewLine("Are u tired of boring console outputs?").Inverse().
		EmptyLine().
		NewLine("✨ Now u can easily create fun console messages! 🦄").Magenta().Bold()

	fmt.Println(message)
}
