local driftwood = require("driftwood")

--- Dumps a table to a string.
---@param tbl table The table to dump.
---@return string The table as a string.
local function dump(tbl)
    if type(tbl) == 'table' then
       local s = '{ '
       for k,v in pairs(tbl) do
          if type(k) ~= 'number' then k = '"'..k..'"' end
          s = s .. '['..k..'] = ' .. dump(v) .. ','
       end
       return s .. '} '
    else
       return tostring(tbl)
    end
end

--- Handles the "selectmenu_demo" application command.
--- @param interaction CommandInteraction The interaction object from Discord.
local function handle_selectmenu(interaction)
    local message_id = driftwood.message.add(interaction.channel_id, "Testing the select menu", {
        driftwood.new_selectmenu("Select an option", "md:selectmenu:initial", {
            driftwood.new_selectmenu_opt("Option 1", "option1"),
            driftwood.new_selectmenu_opt("Option 2", "option2"),
        }, false),
    })

    if not message_id then
        interaction:reply("Failed to send message.", { ephemeral = true })
        return
    end

    interaction:reply("Message sent!", { ephemeral = true })

end

--- Register the /message_demo command.
driftwood.register_application_command({
    name = "seletmenu_demo",
    description = "Demonstrates message management with a select menu.",
    handler = handle_selectmenu
})

--- Register the interaction to handle "md:selectmenu:initial".
driftwood.register_interaction("md:selectmenu:initial", function(interaction)
    -- Check if not null
    if interaction.values == nil then
        interaction:reply("No value found in the interaction.", { ephemeral = true })
        return
    end

    local value = interaction.values[1]
    if value == "option1" then
        interaction:reply("Option 1 selected!", { ephemeral = true })
    elseif value == "option2" then
        interaction:reply("Option 2 selected!", { ephemeral = true })
    else
        interaction:reply("Unknown option selected!" .. dump(interaction.values), { ephemeral = true })
    end
end)