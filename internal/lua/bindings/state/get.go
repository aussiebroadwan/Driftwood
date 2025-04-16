package state

import (
	"driftwood/internal/lua/utils"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// StateBindingGet provides Lua bindings for state management.
type StateBindingGet struct {
	StateManager *utils.StateManager
}

// NewStateBindingGet initializes a new state management instance.
func NewStateBindingGet(sm *utils.StateManager) *StateBindingGet {
	slog.Debug("Creating new StateBindingGet")
	return &StateBindingGet{
		StateManager: sm,
	}
}

// Name returns the name of the binding for global registration in Lua.
func (b *StateBindingGet) Name() string {
	return "get"
}

func (b *StateBindingGet) SetSession(session *discordgo.Session) {}

// Register adds the state-related functions to the Lua state.
func (b *StateBindingGet) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		key := L.CheckString(1)

		value := b.StateManager.Get(key)
		L.Push(value)
		return 1
	})
}

// HandleInteraction is not applicable for this binding.
func (b *StateBindingGet) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *StateBindingGet) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
