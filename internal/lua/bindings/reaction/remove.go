package reaction

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// ReactionBindingRemove provides Lua bindings for Remove reactions to Discord messages.
type ReactionBindingRemove struct {
	Session *discordgo.Session
}

// NewReactionBindingRemove initializes a new reaction Remove instance.
func NewReactionBindingRemove() *ReactionBindingRemove {
	slog.Debug("Creating new ReactionBindingRemove")
	return &ReactionBindingRemove{}
}

// Name returns the name of the binding.
func (b *ReactionBindingRemove) Name() string {
	return "remove"
}

func (b *ReactionBindingRemove) SetSession(session *discordgo.Session) {
	b.Session = session
}

// Register registers the reaction-related functions in the Lua state.
func (b *ReactionBindingRemove) Register() lua.LGFunction {
	return func(L *lua.LState) int {
		messageID := L.CheckString(1)
		channelID := L.CheckString(2)
		content := L.CheckString(3)

		err := b.Session.MessageReactionsRemoveEmoji(channelID, messageID, content)
		if err != nil {
			slog.Error("Failed to react to message", "message_id", messageID, "channel_id", channelID, "error", err)
			L.Push(lua.LFalse)
			return 1
		}

		L.Push(lua.LTrue)
		return 1
	}
}

// HandleInteraction is not applicable for this binding.
func (b *ReactionBindingRemove) HandleInteraction(interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *ReactionBindingRemove) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
