--- Register the /ping_embed command
require("driftwood").register_application_command({
    name = "ping_embed",
    description = "Check bot responsiveness in an embed",
    handler = function(interaction)
        interaction:reply("", {
            ephemeral = true,
            mention = false,
            embed = {
                title = "Pong!",
                description = "The bot is responsive.",
                color = 0x00FF00,
                fields = {
                    {
                        name = "Latency",
                        value = string.format("%dms", 10),
                        inline = true,
                    },
                    {
                        name = "Uptime",
                        value = "1h 23m 45s",
                        inline = false,
                    },
                },
                footer = {
                    text = "This is a footer",
                },
                author = {
                    name = "Bot Name",
                },
            },
        })
    end,
})
