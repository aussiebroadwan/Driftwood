package bindings

import (
	"driftwood/internal/lua/utils"
	"fmt"
	"log/slog"
	"regexp"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// InteractionEventBinding manages custom interaction events in Lua.
type InteractionEventBinding struct {
	Session       *discordgo.Session
	Interactions  map[string]string         // Direct custom_id to Lua handler
	RegexHandlers map[*regexp.Regexp]string // Regex to Lua handler
}

// NewInteractionEventBinding initializes a new InteractionEventBinding instance.
func NewInteractionEventBinding(session *discordgo.Session) *InteractionEventBinding {
	slog.Info("Creating new InteractionEventBinding")
	return &InteractionEventBinding{
		Session:       session,
		Interactions:  make(map[string]string),
		RegexHandlers: make(map[*regexp.Regexp]string),
	}
}

// Name returns the name of the Lua function for this binding.
func (b *InteractionEventBinding) Name() string {
	return "register_interaction"
}

// Register adds the `register_interaction` function to Lua.
func (b *InteractionEventBinding) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		customID := L.CheckString(1)  // First argument is the custom_id or regex pattern
		handler := L.CheckFunction(2) // Second argument is the handler function

		// Create a global function name for the handler
		globalName := fmt.Sprintf("interaction_handler_%s", customID)

		// Set the Lua function as a global
		L.SetGlobal(globalName, handler)

		// Check if the customID is a regex pattern
		if isRegex(customID) {
			compiledRegex, err := regexp.Compile(customID)
			if err != nil {
				L.ArgError(1, fmt.Sprintf("invalid regex pattern: %s", err))
				return 0
			}
			b.RegexHandlers[compiledRegex] = globalName
			slog.Info("Registered regex-based interaction", "pattern", customID, "handler", globalName)
		} else {
			b.Interactions[customID] = globalName
			slog.Info("Registered direct interaction", "custom_id", customID, "handler", globalName)
		}

		return 0
	})
}

func (b *InteractionEventBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return interaction.Type == discordgo.InteractionMessageComponent ||
		interaction.Type == discordgo.InteractionModalSubmit
}

// HandleInteraction executes the Lua handler for a registered interaction event.
func (b *InteractionEventBinding) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	customID := interaction.MessageComponentData().CustomID

	// Check for exact matches first
	if handlerName, exists := b.Interactions[customID]; exists {
		return b.executeHandler(L, interaction, handlerName, customID, nil)
	}

	// Check regex-based handlers
	for pattern, handlerName := range b.RegexHandlers {

		slog.Debug("Checking regex pattern", "pattern", pattern.String(), "custom_id", customID)
		if matches := pattern.FindStringSubmatch(customID); matches != nil {
			groupMap := make(map[string]string)
			for i, name := range pattern.SubexpNames() {
				if i > 0 && name != "" { // Skip the full match and unnamed groups
					groupMap[name] = matches[i]
				}
			}
			return b.executeHandler(L, interaction, handlerName, customID, groupMap)
		}
	}

	slog.Warn("No handler found for interaction", "custom_id", customID)
	return fmt.Errorf("no handler registered for custom ID '%s'", customID)
}

// executeHandler executes the Lua handler for a given custom ID and attaches data from regex matches if available.
func (b *InteractionEventBinding) executeHandler(L *lua.LState, interaction *discordgo.InteractionCreate, handlerName, matchedID string, groupMap map[string]string) error {
	fn := L.GetGlobal(handlerName)
	if fn == lua.LNil {
		slog.Error("Lua handler not implemented", "custom_id", matchedID)
		return fmt.Errorf("handler for custom ID '%s' not implemented", matchedID)
	}

	// Prepare the interaction table
	interactionTable := b.prepareInteractionTable(L, interaction)

	// Add the extracted data from regex as a subtable if available
	if groupMap != nil {
		dataTable := L.NewTable()
		for key, value := range groupMap {
			dataTable.RawSetString(key, lua.LString(value))
		}
		interactionTable.RawSetString("data", dataTable)
	}

	// Call the Lua function
	err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, interactionTable)
	if err != nil {
		slog.Error("Error executing Lua interaction handler", "error", err, "custom_id", matchedID)
		return err
	}

	slog.Info("Interaction handled successfully", "custom_id", matchedID)
	return nil
}

// prepareInteractionTable prepares a Lua table containing interaction details.
func (b *InteractionEventBinding) prepareInteractionTable(L *lua.LState, interaction *discordgo.InteractionCreate) *lua.LTable {
	interactionTable := utils.PrepareInteractionTable(L, b.Session, interaction)
	interactionTable.RawSetString("custom_id", lua.LString(interaction.MessageComponentData().CustomID))
	return interactionTable
}

// isRegex checks if a string is a valid regex pattern.
func isRegex(s string) bool {
	_, err := regexp.Compile(s)
	return err == nil
}
