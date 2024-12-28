local discord = require("driftwood")

-- Define the custom ID for the "Join" button.
local button_id = "example_race_join"

--- Handles the "race" application command.
-- This command sets up a message with an action button to join a snail race.
-- @param interaction table The interaction object from Discord.
local function handle_race_command(interaction)
    -- Respond with a message and a "Join" button.
    interaction:reply_with_action("Join the race!", {
        {
            type = "button",      -- Action type: Button
            label = "Join",       -- Label displayed on the button
            custom_id = button_id, -- Custom ID for interaction tracking
        },
    })
end

--- Handles the "example_race_join" button interaction.
-- This function is triggered when a user clicks the "Join" button.
-- @param interaction table The interaction object from Discord.
local function handle_race_join_interaction(interaction)
    -- Respond with an ephemeral message indicating the user has joined the race.
    interaction:reply("You joined the race!", { ephemeral = true })
end

-- Register the "race" application command.
discord.register_application_command({
    name = "race",               -- Command name
    description = "Manage a snail race", -- Command description
    handler = handle_race_command, -- Link to the handler function
})

-- Register the "example_race_join" interaction for the "Join" button.
discord.register_interaction(button_id, handle_race_join_interaction)
