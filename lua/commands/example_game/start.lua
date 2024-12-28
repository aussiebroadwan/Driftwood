local discord = require("driftwood")

-- Counter variable for tracking delayed messages
local x = 0

--- Sends a delayed message every 5 seconds, up to 5 times.
-- Increments a counter and prints a message. If the counter is less than 5, schedules another message.
local function delayed_message()
    -- Increment the counter
    x = x + 1
    print("This message is displayed after 5 seconds! Count: " .. x)

    -- If the counter is less than 5, schedule another delayed message
    if x < 5 then
        print("Scheduling another message in 5 seconds.")
        discord.run_after(delayed_message, 5) -- Schedule the function to run after 5 seconds
    else
        print("All messages sent!")
    end
end

--- Defines the "start" subcommand.
-- Starts a new game and schedules the first delayed message.
-- @param interaction table The interaction object from Discord.
local function handle_start_command(interaction)
    -- Respond to the user to indicate the game has started
    interaction:reply("Game started!")

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
