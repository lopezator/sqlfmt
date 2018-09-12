package main

import (
	"log"

	"github.com/lopezator/sqlfmt/internal/cli"
)

func main() {
	sqlfmtCmd := cli.NewSqlfmtCmd()

	if err := sqlfmtCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
