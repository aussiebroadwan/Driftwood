package bindings

import (
	"driftwood/internal/lua/utils"
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// MessageBindingEdit provides Lua bindings for managing Discord messages.
type MessageBindingEdit struct {
	Session *discordgo.Session
}

// NewMessageBindingEdit initializes a new message management instance.
func NewMessageBindingEdit(session *discordgo.Session) *MessageBindingEdit {
	slog.Debug("Creating new MessageBindingEdit")
	return &MessageBindingEdit{
		Session: session,
	}
}

// Name returns the name of the binding.
func (b *MessageBindingEdit) Name() string {
	return "edit"
}

// Register registers the message-related functions in the Lua state.
func (b *MessageBindingEdit) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		messageID := L.CheckString(1)
		channelID := L.CheckString(2)
		content := L.CheckString(3)
		components := L.OptTable(4, nil) // Optional components table

		var parsedComponents []discordgo.MessageComponent
		if components != nil {
			var err error
			parsedComponents, err = utils.ParseComponents(L, components)
			if err != nil {
				L.ArgError(4, err.Error())
				return 0
			}
		}

		_, err := b.Session.ChannelMessageEditComplex(&discordgo.MessageEdit{
			ID:         messageID,
			Channel:    channelID,
			Content:    &content,
			Components: &parsedComponents,
		})
		if err != nil {
			slog.Error("Failed to edit message", "message_id", messageID, "channel_id", channelID, "error", err)
			L.Push(lua.LString(fmt.Sprintf("Failed to edit message: %s", err.Error())))
			return 1
		}

		L.Push(lua.LString("Message edited successfully"))
		return 1
	})
}

// HandleInteraction is not applicable for this binding.
func (b *MessageBindingEdit) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *MessageBindingEdit) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
