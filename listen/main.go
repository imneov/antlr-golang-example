package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
)

func executor(in string) {
	expr := parse(in)
	ret := eval(expr)
	fmt.Printf("Answer: %v\n", print(ret))
}

func completer(in prompt.Document) []prompt.Suggest {
	var ret []prompt.Suggest
	return ret
}

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix(">>> "),
		prompt.OptionTitle("calc"),
	)
	p.Run()
}
