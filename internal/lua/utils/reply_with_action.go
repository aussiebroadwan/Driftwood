package utils

import (
	"fmt"
	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// ReplyWithActionFunction returns a Lua function for replying with components (e.g., buttons).
func ReplyWithActionFunction(session *discordgo.Session, interaction *discordgo.InteractionCreate) lua.LGFunction {
	return func(L *lua.LState) int {
		argCount := L.GetTop()
		var content string
		var componentsTable *lua.LTable
		var options *lua.LTable

		if argCount == 4 {
			// Handle the colon syntax (self, content, components, options)
			L.CheckType(1, lua.LTTable) // Ensure the first argument (self) is a table
			content = L.CheckString(2)
			componentsTable = L.CheckTable(3)
			options = L.OptTable(4, nil)
		} else if argCount == 3 {
			// Handle the dot syntax (self, content, components)
			L.CheckType(1, lua.LTTable)
			content = L.CheckString(2)
			componentsTable = L.CheckTable(3)
		} else {
			L.ArgError(1, "invalid arguments, expected (content, components [, options]) or (self, content, components [, options])")
			return 0
		}

		// Parse options
		ephemeral := false
		mention := true
		if options != nil {
			if options.RawGetString("ephemeral") != lua.LNil {
				if options.RawGetString("ephemeral").Type() != lua.LTBool {
					L.ArgError(3, "'ephemeral' in options must be a boolean")
					return 0
				}
				ephemeral = lua.LVAsBool(options.RawGetString("ephemeral"))
			}
			if options.RawGetString("mention") != lua.LNil {
				if options.RawGetString("mention").Type() != lua.LTBool {
					L.ArgError(3, "'mention' in options must be a boolean")
					return 0
				}
				mention = lua.LVAsBool(options.RawGetString("mention"))
			}
		}

		// Modify content to include mention if applicable
		if mention {
			content = fmt.Sprintf("<@%s> %s", interaction.Member.User.ID, content)
		}

		// Set flags for ephemeral messages
		flags := discordgo.MessageFlags(0)
		if ephemeral {
			flags = discordgo.MessageFlagsEphemeral
		}

		// Parse components from the Lua table
		components, err := ParseComponents(L, componentsTable)
		if err != nil {
			L.ArgError(2, err.Error())
			return 0
		}

		// Send the response
		err = session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content:    content,
				Flags:      flags,
				Components: components,
			},
		})
		if err != nil {
			slog.Error("Failed to send interaction reply with components", "error", err)
		}

		return 0
	}
}
