package bindings

import (
	"driftwood/internal/lua/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type NewOptionNumberBinding struct{}

// NewNewOptionNumberBinding initializes a new NewOptionNumberBinding.
func NewNewOptionNumberBinding() *NewOptionNumberBinding {
	slog.Debug("Creating new NewOptionNumberBinding")
	return &NewOptionNumberBinding{}
}

// Name returns the name of the Lua global table for this binding.
func (b *NewOptionNumberBinding) Name() string {
	return "new_number"
}

func (b *NewOptionNumberBinding) SetSession(session *discordgo.Session) {}

// Register adds the `register_application_command` function to a Lua table.
func (b *NewOptionNumberBinding) Register(L *lua.LState) *lua.LFunction {
	slog.Info("Registering new button command Lua function")
	return L.NewFunction(utils.NewOptionRegister(discordgo.ApplicationCommandOptionNumber))
}

// HandleInteraction is not applicable for this binding.
func (b *NewOptionNumberBinding) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *NewOptionNumberBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
