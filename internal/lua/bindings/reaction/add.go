package reaction

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// ReactionBindingAdd provides Lua bindings for add reactions to Discord messages.
type ReactionBindingAdd struct {
	Session *discordgo.Session
}

// NewReactionBindingAdd initializes a new reaction add instance.
func NewReactionBindingAdd() *ReactionBindingAdd {
	slog.Debug("Creating new ReactionBindingAdd")
	return &ReactionBindingAdd{}
}

// Name returns the name of the binding.
func (b *ReactionBindingAdd) Name() string {
	return "add"
}

func (b *ReactionBindingAdd) SetSession(session *discordgo.Session) {
	b.Session = session
}

// Register registers the reaction-related functions in the Lua state.
func (b *ReactionBindingAdd) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		messageID := L.CheckString(1)
		channelID := L.CheckString(2)
		content := L.CheckString(3)

		err := b.Session.MessageReactionAdd(channelID, messageID, content)
		if err != nil {
			slog.Error("Failed to react to message", "message_id", messageID, "channel_id", channelID, "error", err)
			L.Push(lua.LFalse)
			return 1
		}

		L.Push(lua.LTrue)
		return 1
	})
}

// HandleInteraction is not applicable for this binding.
func (b *ReactionBindingAdd) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *ReactionBindingAdd) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
