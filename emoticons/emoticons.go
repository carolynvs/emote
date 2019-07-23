package emoticons

import (
	"fmt"
	"io"
	"os"

	"github.com/atotto/clipboard"
	"github.com/carolynvs/emote/config"
)

type App struct {
	Out    io.Writer
	Config *config.Config
}

func New() (*App, error) {
	c, err := config.Load()
	if err != nil {
		return nil, err
	}
	a := &App{
		Out:    os.Stdout,
		Config: c,
	}
	return a, nil
}

func (a *App) Emote(name string, dest string) {
	emoticon := a.Config.Emoticon[name]

	switch dest {
	case "clipboard":
		clipboard.WriteAll(emoticon)
		fmt.Fprintf(a.Out, emoticon, "was copied to the clipboard")
	default:
		fmt.Fprintln(a.Out, emoticon)
	}
}
