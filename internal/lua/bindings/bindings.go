package bindings

import (
	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type LuaBinding interface {
	Name() string
	Register(L *lua.LState) *lua.LFunction
	HandleCommand(L *lua.LState, interaction *discordgo.InteractionCreate) error
}
