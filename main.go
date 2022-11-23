package main

import (
	"github.com/spf13/cobra"
	"log"
	"test/fare"
)

var command = &cobra.Command{
	Use:     "demo",
	Short:   "demo for shendun",
	Long:    `demo for shendun`,
	Version: "v0.1",
}

func init() {
	command.AddCommand(fare.CmdInit)
	command.AddCommand(fare.CmdCalc)
}

func main() {
	if err := command.Execute(); err != nil {
		log.Fatal(err)
	}
}
