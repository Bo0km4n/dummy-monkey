package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"

	"github.com/Bo0km4n/dummy-monkey/repl"
)

var (
	f = flag.String("file", "", "input file")
)

func init() {
	flag.Parse()
}

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	if *f != "" {
		file, err := os.Open(*f)
		if err != nil {
			panic(err)
		}
		repl.FileExecute(file, os.Stdout)
	} else {
		fmt.Printf("Hello %s! This is the Monket programming language!\n", user.Username)
		fmt.Printf("Feel free to type in commands\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
