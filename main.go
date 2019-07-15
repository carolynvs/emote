package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := buildEmoteCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildEmoteCommand() *cobra.Command {
	emote := &cobra.Command{
		Use:   "emote",
	}
	shrug := &cobra.Command{
		Use: "shrug",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(`¯\_(ツ)_/¯`)
		},
	}
	emote.AddCommand(shrug)
	return emote
}
