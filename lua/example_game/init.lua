local discord = require("driftwood")

-- Import subcommands.
local start_subcommand = require("example_game.start")
local join_subcommand = require("example_game.join")

-- Register the "game" command with subcommands and arguments.
discord.register_application_command({
    name = "game",
    description = "Manage and play games",
    options = {
        start_subcommand,
        join_subcommand,
    },
})