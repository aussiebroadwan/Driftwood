local discord = require("driftwood")

--- Handles the /message_demo command
-- Demonstrates adding, editing, and deleting a message with state management.
-- @param interaction table The interaction object from Discord.
local function handle_message_demo(interaction)
    -- State key for storing message details
    local state_key = "message_demo_state_" .. interaction.interaction_id

    -- Step 1: Add a message to the channel.
    local message_id = discord.message.add(interaction.channel_id, "Initial message! Updating soon...", {
        { type = "button", label = "Click me", custom_id = "md:button:initial" }, 
    })

    if not message_id then
        interaction:reply("Failed to send message.", { ephemeral = true })
        return
    end

    -- Store the message ID and channel ID in the state.
    discord.state.set(state_key, { message_id = message_id, channel_id = interaction.channel_id })

    -- Step 2: Reply with an ephemeral confirmation.
    interaction:reply("Message sent! It will be updated and deleted shortly.", { ephemeral = true })

    discord.log.info("Message added with ID: " .. message_id)

    -- Step 3: Schedule an edit after 5 seconds.
    discord.timer.run_after(function()
        -- Retrieve the stored message details from the state.
        local message_state = discord.state.get(state_key)
        if not message_state then
            discord.log.error("Message state not found. Skipping edit.")
            return
        end

        local success = discord.message.edit(message_state.message_id, message_state.channel_id, "This is the updated content!", {
            { type = "button", label = "Updated button", custom_id = "md:button:updated"}, 
        })

        if not success then
            discord.log.error("Failed to edit the message.")
        end
    end, 5)

    -- Step 4: Schedule a deletion after 10 seconds.
    discord.timer.run_after(function()
        -- Retrieve the stored message details from the state.
        local message_state = discord.state.get(state_key)
        if not message_state then
            discord.log.error("Message state not found. Skipping deletion.")
            return
        end

        local success = discord.message.delete(message_state.message_id, message_state.channel_id)
        if not success then
            discord.log.info("Failed to delete the message.")
        end

        -- Clear the state after the message is deleted.
        discord.state.clear(state_key)
    end, 10)
end

-- Register the /message_demo command.
discord.register_application_command({
    name = "message_demo",
    description = "Demonstrates message lifecycle: add, edit, delete with state management",
    handler = handle_message_demo
})

-- Register the interaction to handle "md:button:value".
discord.register_interaction("md:button:(?P<value>\\d+)", function(interaction)
    local value = interaction.data.value
    if value then
        interaction:reply("Button clicked with value: " .. value, { ephemeral = true })
    else
        interaction:reply("No value found in the interaction.", { ephemeral = true })
    end
end)