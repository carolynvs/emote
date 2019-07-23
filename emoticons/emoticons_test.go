package emoticons

import (
	"bytes"
	"testing"

	"github.com/carolynvs/emote/config"
	"github.com/stretchr/testify/assert"
)

func TestApp_Emote(t *testing.T) {
	const shrugEmoticon = `¯\_(ツ)_/¯`

	out := &bytes.Buffer{}
	app := &App{
		Out: out,
		Config: &config.Config{
			Emoticon: map[string]string{"shrug": shrugEmoticon},
		},
	}

	app.Emote("shrug", "")

	assert.Contains(t, out.String(), shrugEmoticon)
}
