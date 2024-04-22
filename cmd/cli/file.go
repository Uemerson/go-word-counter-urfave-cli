package cli

import (
	"fmt"

	"github.com/Uemerson/go-word-counter-urfave-cli/pkg/word"
	"github.com/urfave/cli/v2"
)

func init() {
	rootCmd.Flags = append(rootCmd.Flags,
		&cli.StringFlag{
			Name:     "file",
			Aliases:  []string{"f"},
			Value:    "",
			Usage:    "file path",
			Required: true,
			Action: func(ctx *cli.Context, s string) error {
				wordCount, err := word.Counter(s)
				if err == nil {
					fmt.Println("total words:", wordCount)
				}
				return err
			},
		},
	)
}
