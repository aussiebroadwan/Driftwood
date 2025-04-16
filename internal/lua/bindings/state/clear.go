package state

import (
	"driftwood/internal/lua/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// StateBindingClear provides Lua bindings for state management.
type StateBindingClear struct {
	StateManager *utils.StateManager
}

// NewStateBindingClear initializes a new state management instance.
func NewStateBindingClear(sm *utils.StateManager) *StateBindingClear {
	slog.Debug("Creating new StateBindingClear")
	return &StateBindingClear{
		StateManager: sm,
	}
}

// Name returns the name of the binding for global registration in Lua.
func (b *StateBindingClear) Name() string {
	return "clear"
}

func (b *StateBindingClear) SetSession(session *discordgo.Session) {}

// Register adds the state-related functions to the Lua state.
func (b *StateBindingClear) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(l *lua.LState) int {
		key := L.CheckString(1)

		b.StateManager.Clear(key)
		return 0
	})
}

// HandleInteraction is not applicable for this binding.
func (b *StateBindingClear) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *StateBindingClear) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
