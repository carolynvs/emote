package main

import (
	"log"
	"os"

	"github.com/carolynvs/emote/emoticons"
	"github.com/spf13/cobra"
)

func main() {
	app, err := emoticons.New()
	if err != nil {
		log.Fatal(err)
	}
	cmd := buildEmoteCommand(app)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildEmoteCommand(app *emoticons.App) *cobra.Command {
	var dest string

	emote := &cobra.Command{
		Use: "emote",
		PreRun: func(cmd *cobra.Command, args []string) {
			app.Out = cmd.OutOrStdout()
		},
		Run: func(cmd *cobra.Command, args []string) {
			emoticonName := args[0]
			app.Emote(emoticonName, dest)
		},
		Args: cobra.ExactArgs(1),
	}

	emote.Flags().StringVar(&dest, "dest", "clipboard", "Where to send your emoticon")

	return emote
}
