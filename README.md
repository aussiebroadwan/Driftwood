# Driftwood

Driftwood is a modular Discord bot framework that uses Lua scripts to define application commands and subcommands. This setup allows developers to extend and customize the bot's functionality easily, making it ideal for multi-purpose and complex Discord bot deployments.

## Features

- **Lua Scripting for Commands**: Commands and subcommands are defined in Lua for maximum flexibility and simplicity.
- **Application Command Registration**: Supports registering commands with detailed options, including subcommands and arguments.
- **Dockerised Deployment**: Easily deploy Driftwood as a containerised application.
- **Extensible Architecture**: Add single-file commands or complex modular commands with an intuitive directory structure.
- **Environment Configurations**: Configure the bot using environment variables.

## Environment Variables

The following environment variables are required:

| Variable | Description |
| --- | --- |
| `DISCORD_TOKEN` | Your Discord bot token. |
| `GUILD_ID` | The ID of the guild (server) to register commands. |

## Development Setup

1. Clone the Repository

```bash
git clone https://github.com/lcox74/driftwood.git
cd driftwood
```

2. Build the Project

```bash
go build -o driftwood ./cmd/driftwood.go
```

3. Run the Bot Locally

Set the required environment variables and run the bot:

```bash
export DISCORD_TOKEN=your_token_here
export GUILD_ID=your_guild_id_here
./driftwood
```

## Creating Commands

### Single-File Command Example

```lua
-- file: lua/commands/ping.lua

local discord = require("discord")

-- Register the /ping command
discord.register_application_command({
    name = "ping",
    description = "Check bot responsiveness",
    handler = function(interaction)
        interaction:reply("Pong!")
    end,
}
```

### Modular Command Example

In this the command entry point is the `init.lua` file.

```lua
-- file: lua/commands/example_game/init.lua

local discord = require("discord")

-- Import subcommands.
local start_subcommand = require("commands.example_game.start")
local join_subcommand = require("commands.example_game.join")

-- Register the "game" command.
discord.register_application_command({
    name = "game",
    description = "Manage and play games",
    options = {
        start_subcommand,
        join_subcommand,
    },
})

-- file: lua/commands/example_game/start.lua

local discord = require("discord")

return {
    name = "start",
    description = "Start a new game",
    type = discord.option_subcommand,
    handler = function(interaction)
        interaction:reply("Game started!")
    end,
}

-- file: lua/commands/example_game/join.lua

local discord = require("discord")

return {
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
        },
    },
    handler = function(interaction)
        local game_id = interaction.options.game_id
        local mention = interaction.options.mention or false

        interaction:reply("Joined game with ID: " .. game_id, { ephemeral = true, mention = mention })
    end,
}

```