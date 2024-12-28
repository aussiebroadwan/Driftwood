local discord = require("driftwood")

local start_subcommand = {
    name = "start",
    description = "Start a new game",
    type = discord.option_subcommand,
    handler = function(interaction)
        interaction:reply("Game started!")
    end,
}

return start_subcommand