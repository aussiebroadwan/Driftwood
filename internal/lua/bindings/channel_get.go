package bindings

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// ChannelBindingGet provides Lua bindings for managing Discord messages.
type ChannelBindingGet struct {
	Session *discordgo.Session
	GuildID string
}

// NewChannelBindingGet initializes a new message management instance.
func NewChannelBindingGet(guildID string) *ChannelBindingGet {
	slog.Debug("Creating new ChannelBindingGet")
	return &ChannelBindingGet{
		GuildID: guildID,
	}
}

// Name returns the name of the binding.
func (b *ChannelBindingGet) Name() string {
	return "get"
}

func (b *ChannelBindingGet) SetSession(session *discordgo.Session) {
	b.Session = session
}

// Register registers the channel-related functions in the Lua state.
func (b *ChannelBindingGet) Register(L *lua.LState) *lua.LFunction {
	return L.NewFunction(func(L *lua.LState) int {
		channelName := L.CheckString(1)

		channels, err := b.Session.GuildChannels(b.GuildID)
		if err != nil {
			slog.Error("Failed to get channels", "guild_id", b.GuildID, "error", err)
			L.RaiseError("Failed to get channels: %s", err.Error())
			L.Push(lua.LNil)
			return 1
		}

		for _, channel := range channels {
			if channel.Name == channelName {
				L.Push(lua.LString(channel.ID))
				return 1
			}
		}

		L.Push(lua.LNil)
		return 1
	})
}

func (b *ChannelBindingGet) HandleInteraction(L *lua.LState, interaction *discordgo.InteractionCreate) error {
	return nil
}

func (b *ChannelBindingGet) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
