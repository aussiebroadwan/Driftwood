package bindings

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type NewButtonBinding struct{}

// NewNewButtonBinding initializes a new NewButtonBinding.
func NewNewButtonBinding() *NewButtonBinding {
	return &NewButtonBinding{}
}

// Name returns the name of the Lua global table for this binding.
func (b *NewButtonBinding) Name() string {
	return "new_button"
}

func (b *NewButtonBinding) SetSession(session *discordgo.Session) {}

func (b *NewButtonBinding) Register() lua.LGFunction {
	slog.Info("Registering new button command Lua function")
	return func(L *lua.LState) int {
		argCount := L.GetTop()
		buttonTable := L.NewTable()

		switch argCount {
		case 3:
			buttonTable.RawSetString("label", lua.LString(L.CheckString(1)))
			buttonTable.RawSetString("custom_id", lua.LString(L.CheckString(2)))
			buttonTable.RawSetString("disabled", lua.LBool(L.CheckBool(3)))
		case 2:
			buttonTable.RawSetString("label", lua.LString(L.CheckString(1)))
			buttonTable.RawSetString("custom_id", lua.LString(L.CheckString(2)))
			buttonTable.RawSetString("disabled", lua.LFalse)
		default:
			L.ArgError(1, "invalid arguments, expected (name, custom_id [, disabled])")
		}

		// Create a Table for the button
		buttonTable.RawSetString("type", lua.LString("button"))

		L.Push(buttonTable)
		return 1
	}
}

// HandleInteraction is not applicable for this binding.
func (b *NewButtonBinding) HandleInteraction(interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *NewButtonBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
