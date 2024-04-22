package cli

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var rootCmd = &cli.App{
	Name:  "word-counter",
	Usage: "A straightforward command-line interface (CLI) that counts the number of words in a text file provided as an argument",
	Action: func(*cli.Context) error {
		return nil
	},
}

func Execute() {
	if err := rootCmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
