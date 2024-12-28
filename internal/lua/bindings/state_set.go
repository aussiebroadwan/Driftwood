package bindings

import (
	"driftwood/internal/lua/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// StateBindingSet provides Lua bindings for state management.
type StateBindingSet struct {
	StateManager *utils.StateManager
}

// NewStateBindingSet initializes a new state management instance.
func NewStateBindingSet(sm *utils.StateManager) *StateBindingSet {
	slog.Info("Creating new StateBindingSet")
	return &StateBindingSet{
		StateManager: sm,
	}
}

// Name returns the name of the binding for global registration in Lua.
func (b *StateBindingSet) Name() string {
	return "set"
}

// Register adds the state-related functions to the Lua state.
func (b *StateBindingSet) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {

		key := L.CheckString(1)
		value := L.CheckAny(2)
		expiry := L.OptInt(3, 0) // Optional expiry in seconds

		b.StateManager.Set(key, value, expiry)
		return 0
	})
}

// HandleInteraction is not applicable for this binding.
func (b *StateBindingSet) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *StateBindingSet) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
