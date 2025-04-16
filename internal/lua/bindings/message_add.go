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
func NewMessageBindingAdd() *MessageBindingAdd {
	slog.Debug("Creating new MessageBindingAdd")
	return &MessageBindingAdd{}
}

// Name returns the name of the binding.
func (b *MessageBindingAdd) Name() string {
	return "add"
}

func (b *MessageBindingAdd) SetSession(session *discordgo.Session) {
	b.Session = session
}

// Register registers the message-related functions in the Lua state.
func (b *MessageBindingAdd) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		channelID := L.CheckString(1)
		content := L.CheckString(2)
		opts := L.OptTable(3, nil)

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

		// Parse the embed table if provided
		var parsedComponents []discordgo.MessageComponent
		if components != nil {
			var err error
			parsedComponents, err = utils.ParseComponents(L, components)
			if err != nil {
				L.ArgError(3, err.Error())
				return 0
			}
		}

		// Parse embed if provided.
		var embed *discordgo.MessageEmbed
		if embedTable != nil {
			var err error
			embed, err = utils.ParseEmbed(L, embedTable)
			if err != nil {
				L.ArgError(4, err.Error())
				return 0
			}
		}

		slog.Info("Sending complex message", "channel_id", channelID, "content", content, "components", parsedComponents, "embed", embed)

		message, err := b.Session.ChannelMessageSendComplex(channelID, &discordgo.MessageSend{
			Content:    content,
			Components: parsedComponents,
			Embed:      embed,
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
