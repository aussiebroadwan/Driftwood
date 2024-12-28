# Driftwood

Driftwood is a modular Discord bot framework designed for flexibility and ease of use. By leveraging Lua scripts, it empowers developers to define application commands and subcommands seamlessly, making it an excellent choice for building multi-purpose, customisable Discord bots.

## Features

- **Lua Scripting for Commands**: Commands and subcommands are defined in Lua for maximum flexibility and simplicity.
- **Application Command Registration**: Supports registering commands with detailed options, including subcommands and arguments.
- **Dockerised Deployment**: Easily deploy Driftwood as a containerised application.
- **Extensible Architecture**: Add single-file commands or complex modular commands with an intuitive directory structure.
- **Environment Configurations**: Configure the bot using environment variables.

## Getting Started

To quickly create your own Discord bot using Driftwood, you can use the prebuilt Docker image available on GitHub Container Registry.

### Step 1: Pull the Docker Image

```bash
docker pull ghcr.io/lcox74/driftwood:latest
```

### Step 2: Prepare Your Lua Scripts

Set up a directory for your Lua scripts on your host machine (e.g., `./lua`). Add command scripts to this directory following the structure explained in the Creating Commands section.

### Step 3: Run the Docker Container

Run the container with the required environment variables and mount your Lua script directory:

```bash
docker run -d \
  -e DISCORD_TOKEN=your_token_here \
  -e GUILD_ID=your_guild_id_here \
  -v $(pwd)/lua:/lua \
  ghcr.io/lcox74/driftwood:latest
```

Alternatively, simplify your setup using `docker-compose.yml`.

## Development Setup

### Step 1: Clone the Repository

```bash
git clone https://github.com/lcox74/driftwood.git
cd driftwood
```

### Step 2: Build the Project

```bash
go build -o driftwood ./cmd/driftwood.go
```

### Step 3: Run the Bot Locally

Set the necessary environment variables and run the bot:

```bash
export DISCORD_TOKEN=your_token_here
export GUILD_ID=your_guild_id_here
./driftwood
```

## Environment Variables

The following environment variables are required:

| Variable | Description |
| --- | --- |
| `DISCORD_TOKEN` | Your Discord bot token. |
| `GUILD_ID` | The ID of the guild (server) to register commands. |

## Creating Commands

Driftwood supports both single-file and modular command structures.

### Single-File Command Example

```lua
-- file: lua/commands/ping.lua

local discord = require("driftwood")

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

**Example: Command Entry Point (`init.lua`)**
```lua
-- file: lua/commands/example_game/init.lua

local discord = require("driftwood")

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

```

**Example: Start Subcommand**

```lua
-- file: lua/commands/example_game/start.lua

local discord = require("driftwood")

return {
    name = "start",
    description = "Start a new game",
    type = discord.option_subcommand,
    handler = function(interaction)
        interaction:reply("Game started!")
    end,
}
```

**Example: Join Subcommand**
```lua
-- file: lua/commands/example_game/join.lua

local discord = require("driftwood")

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