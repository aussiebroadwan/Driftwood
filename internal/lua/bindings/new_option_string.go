package bindings

import (
	"driftwood/internal/lua/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

type NewOptionStringBinding struct{}

// NewNewOptionStringBinding initializes a new NewOptionStringBinding.
func NewNewOptionStringBinding() *NewOptionStringBinding {
	slog.Debug("Creating new RunAfterBinding")
	return &NewOptionStringBinding{}
}

// Name returns the name of the Lua global table for this binding.
func (b *NewOptionStringBinding) Name() string {
	return "new_string"
}

func (b *NewOptionStringBinding) SetSession(session *discordgo.Session) {}

// Register adds the `register_application_command` function to a Lua table.
func (b *NewOptionStringBinding) Register(L *lua.LState) *lua.LFunction {
	slog.Info("Registering new button command Lua function")
	return L.NewFunction(utils.NewOptionRegister(discordgo.ApplicationCommandOptionString))
}

// HandleInteraction is not applicable for this binding.
func (b *NewOptionStringBinding) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *NewOptionStringBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
