--- Discord module for managing bot commands and interactions.
-- @module discord

local discord = {}

--- Registers an application command with Discord.
-- @param command table A table defining the command.
-- @field name string The name of the command (required).
-- @field description string The description of the command (required).
-- @field options table|nil A table defining the command's options (optional).
-- @field options[].name string The name of the option (required for each option).
-- @field options[].description string The description of the option (required for each option).
-- @field options[].type number The type of the option, based on Discord's ApplicationCommandOptionType enum (required).
-- @field options[].required boolean Whether the option is mandatory (optional, defaults to false).
-- @field options[].options table|nil Sub-options for subcommands or groups (optional).
-- @field handler function The Lua function to handle the command (required).
-- The handler function receives an `interaction` object with a `reply` method.
-- @usage
-- discord.register_application_command({
--     name = "game",
--     description = "Manage and play games",
--     options = {
--         {
--             name = "start",
--             description = "Start a new game",
--             type = 1, -- Subcommand
--             options = {
--                 {
--                     name = "team",
--                     description = "Name of the team",
--                     type = 3, -- String
--                     required = true
--                 }
--             },
--             handler = function(interaction)
--                 interaction:reply("Game started!")
--             end
--         },
--         {
--             name = "join",
--             description = "Join an existing game",
--             type = 1, -- Subcommand
--             options = {
--                 {
--                     name = "game_id",
--                     description = "ID of the game to join",
--                     type = 3, -- String
--                     required = true
--                 }
--             },
--             handler = function(interaction)
--                 local game_id = interaction.options.game_id
--                 interaction:reply("Joined the game with ID: " .. game_id)
--             end
--         }
--     },
-- })
function discord.register_application_command(command)
    -- Placeholder for IDE type hinting. Actual implementation in Go.
end

return discord
