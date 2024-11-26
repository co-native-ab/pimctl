package main

import (
	"log"

	"github.com/co-native-ab/pimctl/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cmd.RootCmd, "./generated-docs")
	if err != nil {
		log.Fatal(err)
	}
}
