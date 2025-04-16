package options

import (
	"driftwood/internal/lua/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type NewOptionBoolBinding struct{}

// NewNewOptionBoolBinding initializes a new NewOptionBoolBinding.
func NewNewOptionBoolBinding() *NewOptionBoolBinding {
	slog.Debug("Creating new NewOptionBoolBinding")
	return &NewOptionBoolBinding{}
}

// Name returns the name of the Lua global table for this binding.
func (b *NewOptionBoolBinding) Name() string {
	return "new_bool"
}

func (b *NewOptionBoolBinding) SetSession(session *discordgo.Session) {}

// Register adds the `register_application_command` function to a Lua table.
func (b *NewOptionBoolBinding) Register(L *lua.LState) *lua.LFunction {
	slog.Info("Registering new button command Lua function")
	return L.NewFunction(utils.NewOptionRegister(discordgo.ApplicationCommandOptionBoolean))
}

// HandleInteraction is not applicable for this binding.
func (b *NewOptionBoolBinding) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *NewOptionBoolBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
