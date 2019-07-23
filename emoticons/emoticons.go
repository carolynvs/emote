package emoticons

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/spf13/viper"
)

type App struct {
	viper *viper.Viper
}

func New() (*App, error) {
	v := viper.New()
	v.SetConfigName("emote")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &App{viper: v}, nil
}

func (a *App) Emote(name string, dest string) {
	emoticon := a.viper.GetString("emoticon."+name)

	switch dest {
	case "clipboard":
		clipboard.WriteAll(emoticon)
		fmt.Println(emoticon, "was copied to the clipboard")
	default:
		fmt.Println(emoticon)
	}
}
