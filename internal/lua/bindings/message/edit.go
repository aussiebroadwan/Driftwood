package message

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
func NewMessageBindingEdit() *MessageBindingEdit {
	slog.Debug("Creating new MessageBindingEdit")
	return &MessageBindingEdit{}
}

// Name returns the name of the binding.
func (b *MessageBindingEdit) Name() string {
	return "edit"
}

func (b *MessageBindingEdit) SetSession(session *discordgo.Session) {
	b.Session = session
}

// Register registers the message-related functions in the Lua state.
func (b *MessageBindingEdit) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		messageID := L.CheckString(1)
		channelID := L.CheckString(2)
		content := L.CheckString(3)
		opts := L.OptTable(4, nil)

		// Options table may have "components" and "embed" keys
		var components *lua.LTable = nil
		var embedTable *lua.LTable = nil

		if opts != nil {
			comp := opts.RawGetString("components")
			if comp != lua.LNil {
				if compTable, ok := comp.(*lua.LTable); ok {
					components = compTable
				} else {
					L.ArgError(3, "options.components must be a table")
					return 0
				}
			}

			em := opts.RawGetString("embed")
			if em != lua.LNil {
				if embedT, ok := em.(*lua.LTable); ok {
					embedTable = embedT
				} else {
					L.ArgError(3, "options.embed must be a table")
					return 0
				}
			}
		}

		var parsedComponents []discordgo.MessageComponent
		if components != nil {
			var err error
			parsedComponents, err = utils.ParseComponents(L, components)
			if err != nil {
				L.ArgError(4, err.Error())
				return 0
			}
		}

		// Parse embed if provided.
		var embed *discordgo.MessageEmbed
		if embedTable != nil {
			var err error
			embed, err = utils.ParseEmbed(L, embedTable)
			if err != nil {
				L.ArgError(5, err.Error())
				return 0
			}
		}

		_, err := b.Session.ChannelMessageEditComplex(&discordgo.MessageEdit{
			ID:         messageID,
			Channel:    channelID,
			Content:    &content,
			Components: &parsedComponents,
			Embed:      embed,
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
