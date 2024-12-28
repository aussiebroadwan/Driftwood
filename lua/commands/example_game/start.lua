local discord = require("driftwood")

--- Sends a delayed message every 5 seconds, up to 5 times.
-- Increments a counter and prints a message. If the counter is less than 5, schedules another message.
local function delayed_message()
    data = discord.state_get("game_data") -- Retrieve game data from the state

    -- Check if the game data is missing
    if data == nil then
        print("Game data is missing!")
        return
    end

    -- Increment the counter
    data.x = data.x + 1
    print("This message is displayed after 5 seconds! Count: " .. data.x)

    -- If the counter is less than 5, schedule another delayed message
    if data.x < 5 then
        print("Scheduling another message in 5 seconds.")
        discord.state_set("game_data", data) -- Store the updated counter in the state
        discord.run_after(delayed_message, 5) -- Schedule the function to run after 5 seconds
    else
        print("All messages sent!")
        discord.state_clear("game_data") -- Clear the game data from the state
    end
end

--- Defines the "start" subcommand.
-- Starts a new game and schedules the first delayed message.
-- @param interaction table The interaction object from Discord.
local function handle_start_command(interaction)
    -- Respond to the user to indicate the game has started
    interaction:reply("Game started!")
    discord.state_set("game_data", { x = 0 }) -- Store game data in the state

    -- Schedule the first delayed message to run after 5 seconds
    discord.run_after(delayed_message, 5)
end

-- Define the "start" subcommand metadata
local start_subcommand = {
    name = "start",                     -- Subcommand name
    description = "Start a new game",   -- Subcommand description
    type = discord.option_subcommand,   -- Subcommand type
    handler = handle_start_command,     -- Link to the handler function
}

return start_subcommand
