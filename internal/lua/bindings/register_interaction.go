package bindings

import (
	"driftwood/internal/lua/utils"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// InteractionEventBinding manages custom interaction events in Lua.
type InteractionEventBinding struct {
	Session      *discordgo.Session
	Interactions map[string]string // Maps custom_id to Lua global handler names
}

// NewInteractionEventBinding initializes a new InteractionEventBinding instance.
func NewInteractionEventBinding(session *discordgo.Session) *InteractionEventBinding {
	slog.Info("Creating new InteractionEventBinding")
	return &InteractionEventBinding{
		Session:      session,
		Interactions: make(map[string]string),
	}
}

// Name returns the name of the Lua function for this binding.
func (b *InteractionEventBinding) Name() string {
	return "register_interaction"
}

// Register adds the `register_interaction` function to Lua.
func (b *InteractionEventBinding) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		// Retrieve the custom ID and handler from the arguments
		customID := L.CheckString(1)  // First argument is the custom_id
		handler := L.CheckFunction(2) // Second argument is the handler function

		// Create a global function name for the handler
		globalName := fmt.Sprintf("interaction_handler_%s", customID)

		// Set the Lua function as a global
		L.SetGlobal(globalName, handler)

		// Store the mapping for future interaction handling
		b.Interactions[customID] = globalName

		slog.Info("Registered interaction", "custom_id", customID, "handler", globalName)
		return 0
	})
}

func (b *InteractionEventBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return interaction.Type == discordgo.InteractionMessageComponent ||
		interaction.Type == discordgo.InteractionModalSubmit
}

// HandleInteraction executes the Lua handler for a registered interaction event.
func (b *InteractionEventBinding) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// Get the custom ID from the interaction
	customID := interaction.MessageComponentData().CustomID

	// Look up the Lua function for the custom ID
	handlerName, exists := b.Interactions[customID]
	if !exists {
		slog.Warn("Custom ID not registered", "custom_id", customID)
		return fmt.Errorf("custom ID '%s' not registered", customID)
	}

	// Retrieve the Lua function
	fn := L.GetGlobal(handlerName)
	if fn == lua.LNil {
		slog.Error("Lua handler not implemented", "custom_id", customID)
		return fmt.Errorf("handler for custom ID '%s' not implemented", customID)
	}

	// Create a Lua table representing the interaction
	interactionTable := b.prepareInteractionTable(L, interaction)

	// Call the Lua function
	err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, interactionTable)

	if err != nil {
		slog.Error("Error executing Lua interaction handler", "error", err, "custom_id", customID)
		return err
	}

	slog.Info("Interaction handled successfully", "custom_id", customID)
	return nil
}

// prepareInteractionTable prepares a Lua table containing interaction details.
func (b *InteractionEventBinding) prepareInteractionTable(L *lua.LState, interaction *discordgo.InteractionCreate) *lua.LTable {
	interactionTable := L.NewTable()

	// Add interaction data (e.g., custom_id) to the table
	interactionTable.RawSetString("custom_id", lua.LString(interaction.MessageComponentData().CustomID))

	// Add the `reply` method to the interaction table
	interactionTable.RawSetString("reply", L.NewFunction(b.replyFunction(interaction)))
	interactionTable.RawSetString("reply_with_action", L.NewFunction(b.replyWithActionFunction(interaction)))

	return interactionTable
}

// replyFunction returns a Lua function for replying to interactions.
func (b *InteractionEventBinding) replyFunction(interaction *discordgo.InteractionCreate) lua.LGFunction {
	return utils.ReplyFunction(b.Session, interaction)
}

// replyWithActionFunction returns a Lua function for replying to interactions.
func (b *InteractionEventBinding) replyWithActionFunction(interaction *discordgo.InteractionCreate) lua.LGFunction {
	return utils.ReplyWithActionFunction(b.Session, interaction)
}
