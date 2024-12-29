local driftwood = require("driftwood")

--- Handles the "join" subcommand with arguments and options.
--- @param interaction CommandInteraction The interaction object from Discord.
local function handle_join(interaction)
    -- Retrieve arguments from the interaction.
    local game_id = interaction.options.game_id
    local mention = interaction.options.mention or false -- Default to false if not provided.

    -- Build the response.
    local response = "Joined the game with ID: `" .. game_id .. "`"

    -- Send a reply with options.
    interaction:reply(response, { ephemeral = true, mention = mention })
end

--- Define the "join" subcommand metadata.
--- @type CommandOption
local join_command =  {
    name = "join",
    description = "Join an existing game",
    type = driftwood.option_subcommand,
    options = {
        driftwood.option.new_string("game_id", "ID of the game to join", true),
        driftwood.option.new_bool("mention", "Mention the user in the response"),
    },
    handler = handle_join,
}

return join_command