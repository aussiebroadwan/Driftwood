package utils

import (
	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// PrepareInteractionTable prepares a Lua table containing interaction details.
func PrepareInteractionTable(L *lua.LState, session *discordgo.Session, interaction *discordgo.InteractionCreate) *lua.LTable {
	interactionTable := L.NewTable()

	// Add the `reply` method to the interaction table
	interactionTable.RawSetString("reply", L.NewFunction(ReplyFunction(session, interaction)))
	interactionTable.RawSetString("reply_with_action", L.NewFunction(ReplyWithActionFunction(session, interaction)))

	interactionTable.RawSetString("interaction_id", lua.LString(interaction.ID))
	interactionTable.RawSetString("channel_id", lua.LString(interaction.ChannelID))

	return interactionTable
}
