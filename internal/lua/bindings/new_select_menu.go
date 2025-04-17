package bindings

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type NewSelectMenuBinding struct{}

// NewNewSelectMenuBinding initializes a new NewSelectMenuBinding.
func NewNewSelectMenuBinding() *NewSelectMenuBinding {
	return &NewSelectMenuBinding{}
}

// Name returns the name of the Lua global table for this binding.
func (b *NewSelectMenuBinding) Name() string {
	return "new_selectmenu"
}

func (b *NewSelectMenuBinding) SetSession(session *discordgo.Session) {}

func (b *NewSelectMenuBinding) Register() lua.LGFunction {
	slog.Info("Registering new select menu command Lua function")
	return func(L *lua.LState) int {
		argCount := L.GetTop()
		selectTable := L.NewTable()

		switch argCount {
		case 4:
			selectTable.RawSetString("placeholder", lua.LString(L.CheckString(1)))
			selectTable.RawSetString("custom_id", lua.LString(L.CheckString(2)))
			selectTable.RawSetString("options", L.CheckTable(3))
			selectTable.RawSetString("disabled", lua.LBool(L.CheckBool(4)))
		case 3:
			selectTable.RawSetString("placeholder", lua.LString(L.CheckString(1)))
			selectTable.RawSetString("custom_id", lua.LString(L.CheckString(2)))
			selectTable.RawSetString("disabled", lua.LFalse)
			selectTable.RawSetString("options", L.CheckTable(3))
		default:
			L.ArgError(1, "invalid arguments, expected (name, custom_id [, disabled])")
		}

		// Create a Table for the button
		selectTable.RawSetString("type", lua.LString("select"))

		L.Push(selectTable)
		return 1
	}
}

// HandleInteraction is not applicable for this binding.
func (b *NewSelectMenuBinding) HandleInteraction(interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *NewSelectMenuBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
