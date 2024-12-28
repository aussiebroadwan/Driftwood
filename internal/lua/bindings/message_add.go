package bindings

import (
	"driftwood/internal/lua/utils"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// MessageBindingAdd provides Lua bindings for managing Discord messages.
type MessageBindingAdd struct {
	Session *discordgo.Session
}

// NewMessageBindingAdd initializes a new message management instance.
func NewMessageBindingAdd(session *discordgo.Session) *MessageBindingAdd {
	slog.Info("Creating new MessageBindingAdd")
	return &MessageBindingAdd{
		Session: session,
	}
}

// Name returns the name of the binding.
func (b *MessageBindingAdd) Name() string {
	return "add_message"
}

// Register registers the message-related functions in the Lua state.
func (b *MessageBindingAdd) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		channelID := L.CheckString(1)
		content := L.CheckString(2)
		components := L.OptTable(3, nil) // Optional components table

		var parsedComponents []discordgo.MessageComponent
		if components != nil {
			var err error
			parsedComponents, err = utils.ParseComponents(L, components)
			if err != nil {
				L.ArgError(3, err.Error())
				return 0
			}
		}

		message, err := b.Session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
			Content:    content,
			Components: parsedComponents,
		})
		if err != nil {
			slog.Error("Failed to send message", "channel_id", channelID, "error", err)
			L.Push(lua.LNil)
			L.Push(lua.LString(fmt.Sprintf("Failed to send message: %s", err.Error())))
			return 2
		}

		L.Push(lua.LString(message.ID)) // Return the message ID
		return 1
	})
}

// HandleInteraction is not applicable for this binding.
func (b *MessageBindingAdd) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *MessageBindingAdd) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
