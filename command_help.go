package main

import (
	"fmt"
	"io"
)

func displayHelp(out io.Writer) error {
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Welcome to the Pokedex!\nUsage:")
	fmt.Fprintln(out)
	for _, command := range getCommands() {
		fmt.Fprintf(out, "%v: %v\n", command.name, command.description)
	}
	return nil
}
