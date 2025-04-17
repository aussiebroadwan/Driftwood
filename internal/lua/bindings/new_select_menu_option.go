package bindings

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type NewSelectMenuOptionBinding struct{}

// NewNewSelectMenuOptionBinding initializes a new NewSelectMenuOptionBinding.
func NewNewSelectMenuOptionBinding() *NewSelectMenuOptionBinding {
	return &NewSelectMenuOptionBinding{}
}

// Name returns the name of the Lua global table for this binding.
func (b *NewSelectMenuOptionBinding) Name() string {
	return "new_selectmenu_opt"
}

func (b *NewSelectMenuOptionBinding) SetSession(session *discordgo.Session) {}

func (b *NewSelectMenuOptionBinding) Register() lua.LGFunction {
	slog.Info("Registering new select menu command Lua function")
	return func(L *lua.LState) int {
		optTable := L.NewTable()
		optTable.RawSetString("label", lua.LString(L.CheckString(1)))
		optTable.RawSetString("value", lua.LString(L.CheckString(2)))
		L.Push(optTable)
		return 1
	}
}

// HandleInteraction is not applicable for this binding.
func (b *NewSelectMenuOptionBinding) HandleInteraction(interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *NewSelectMenuOptionBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
