local driftwood = require("driftwood")

--- Register a handler for when the bot is ready.
driftwood.on_ready("first_ready", function ()

    -- Get the ID of the "general" channel.
    local channel_id = driftwood.channel.get("general")
    if channel_id == nil then
        driftwood.log.error("Channel not found")
        return
    end

    -- Add a message to the channel.
    driftwood.message.add(channel_id, "Hello, world!")
end)