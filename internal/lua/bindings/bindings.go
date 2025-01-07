package bindings

import (
	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type LuaBinding interface {
	Name() string
	Register(L *lua.LState) *lua.LFunction
	SetSession(session *discordgo.Session)
	HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error
	CanHandleInteraction(interaction *discordgo.InteractionCreate) bool
}
