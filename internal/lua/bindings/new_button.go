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

func (b *NewButtonBinding) Register(L *lua.LState) *lua.LFunction {
	slog.Info("Registering new button command Lua function")
	return L.NewFunction(func(L *lua.LState) int {
		argCount := L.GetTop()
		buttonTable := L.NewTable()

		if argCount == 2 {
			buttonTable.RawSetString("label", lua.LString(L.CheckString(1)))
			buttonTable.RawSetString("custom_id", lua.LString(L.CheckString(2)))
			buttonTable.RawSetString("disabled", lua.LFalse)
		} else if argCount == 3 {
			buttonTable.RawSetString("label", lua.LString(L.CheckString(1)))
			buttonTable.RawSetString("custom_id", lua.LString(L.CheckString(2)))
			buttonTable.RawSetString("disabled", lua.LBool(L.CheckBool(3)))
		} else {
			L.ArgError(1, "invalid arguments, expected (name, custom_id [, disabled])")
		}

		// Create a Table for the button
		buttonTable.RawSetString("type", lua.LString("button"))

		L.Push(buttonTable)
		return 1
	})
}

// HandleInteraction is not applicable for this binding.
func (b *NewButtonBinding) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *NewButtonBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
