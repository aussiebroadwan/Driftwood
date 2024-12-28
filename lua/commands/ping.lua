
-- Register the /ping command
require("discord").register_application_command({
    name = "ping",
    description = "Check bot responsiveness",
    handler = function(interaction)
        interaction:reply("Pong!")
    end,
})
