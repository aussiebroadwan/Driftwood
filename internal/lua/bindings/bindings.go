package bindings

import (
	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type LuaBinding interface {
	Name() string
	Register() lua.LGFunction
	SetSession(session *discordgo.Session)
	HandleInteraction(interaction *discordgo.InteractionCreate) error
	CanHandleInteraction(interaction *discordgo.InteractionCreate) bool
}
