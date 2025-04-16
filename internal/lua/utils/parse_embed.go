package utils

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// ParseEmbed parses a Lua table into a Discord message embed.
func ParseEmbed(L *lua.LState, table *lua.LTable) (*discordgo.MessageEmbed, error) {
	embed := &discordgo.MessageEmbed{}

	// Title
	title := table.RawGetString("title")
	if title != lua.LNil {
		embed.Title = title.String()
	}

	// Description
	description := table.RawGetString("description")
	if description != lua.LNil {
		embed.Description = description.String()
	}

	// URL
	url := table.RawGetString("url")
	if url != lua.LNil {
		embed.URL = url.String()
	}

	// Color (expects a number)
	color := table.RawGetString("color")
	if color != lua.LNil {
		if color.Type() != lua.LTNumber {
			return nil, fmt.Errorf("embed.color must be a number")
		}
		embed.Color = int(color.(lua.LNumber))
	}

	// Image (expects a table with key "url")
	imageRaw := table.RawGetString("image")
	if imageRaw != lua.LNil {
		if imageTable, ok := imageRaw.(*lua.LTable); ok {
			imgURL := imageTable.RawGetString("url")
			if imgURL != lua.LNil {
				embed.Image = &discordgo.MessageEmbedImage{
					URL: imgURL.String(),
				}
			}
		}
	}

	// Thumbnail (expects a table with key "url")
	thumbnailRaw := table.RawGetString("thumbnail")
	if thumbnailRaw != lua.LNil {
		if thumbnailTable, ok := thumbnailRaw.(*lua.LTable); ok {
			thumbURL := thumbnailTable.RawGetString("url")
			if thumbURL != lua.LNil {
				embed.Thumbnail = &discordgo.MessageEmbedThumbnail{
					URL: thumbURL.String(),
				}
			}
		}
	}

	// Footer (expects a table with keys "text" and optional "icon_url")
	footerRaw := table.RawGetString("footer")
	if footerRaw != lua.LNil {
		if footerTable, ok := footerRaw.(*lua.LTable); ok {
			footerText := footerTable.RawGetString("text")
			if footerText != lua.LNil {
				footer := discordgo.MessageEmbedFooter{
					Text: footerText.String(),
				}
				iconURL := footerTable.RawGetString("icon_url")
				if iconURL != lua.LNil {
					footer.IconURL = iconURL.String()
				}
				embed.Footer = &footer
			}
		}
	}

	// Author (expects a table with keys "name", "url", and "icon_url")
	authorRaw := table.RawGetString("author")
	if authorRaw != lua.LNil {
		if authorTable, ok := authorRaw.(*lua.LTable); ok {
			authorName := authorTable.RawGetString("name")
			if authorName != lua.LNil {
				author := discordgo.MessageEmbedAuthor{
					Name: authorName.String(),
				}
				authorURL := authorTable.RawGetString("url")
				if authorURL != lua.LNil {
					author.URL = authorURL.String()
				}
				iconURL := authorTable.RawGetString("icon_url")
				if iconURL != lua.LNil {
					author.IconURL = iconURL.String()
				}
				embed.Author = &author
			}
		}
	}

	// Fields (expects an array/table of field tables with keys "name", "value", and optional "inline")
	fieldsRaw := table.RawGetString("fields")
	if fieldsRaw != lua.LNil {
		if fieldsTable, ok := fieldsRaw.(*lua.LTable); ok {
			var fields []*discordgo.MessageEmbedField
			fieldsTable.ForEach(func(_, value lua.LValue) {
				if fieldTable, ok := value.(*lua.LTable); ok {
					field := &discordgo.MessageEmbedField{}
					name := fieldTable.RawGetString("name")
					if name != lua.LNil {
						field.Name = name.String()
					}
					val := fieldTable.RawGetString("value")
					if val != lua.LNil {
						field.Value = val.String()
					}
					inlineRaw := fieldTable.RawGetString("inline")
					if inlineRaw.Type() == lua.LTBool {
						field.Inline = lua.LVAsBool(inlineRaw)
					}
					fields = append(fields, field)
				}
			})
			embed.Fields = fields
		}
	}

	return embed, nil
}
