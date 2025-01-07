local driftwood = require("driftwood")

--- Register a handler for when the bot is ready.
driftwood.on_ready("first_ready", function ()
    driftwood.message.add("1323233843444060220", "Hello, world!")
end)