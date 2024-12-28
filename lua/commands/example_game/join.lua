local discord = require("discord")

--- Handles the "join" subcommand with arguments and options.
-- @param interaction table The interaction object from Discord.
local function handle_join(interaction)
    -- Retrieve arguments from the interaction.
    local game_id = interaction.options.game_id
    local mention = interaction.options.mention or false -- Default to false if not provided.

    -- Build the response.
    local response = "Joined the game with ID: `" .. game_id .. "`"

    -- Send a reply with options.
    interaction:reply(response, { ephemeral = true, mention = mention })
end

local join_command =  {
    name = "join",
    description = "Join an existing game",
    type = discord.option_subcommand,
    options = {
        {
            name = "game_id",
            description = "ID of the game to join",
            type = discord.option_string,
            required = true,
        },
        {
            name = "mention",
            description = "Mention the user in the response",
            type = discord.option_boolean,
            required = false,
        }
    },
    handler = handle_join,
}

return join_command