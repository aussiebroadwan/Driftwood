local driftwood = require("driftwood")

--- Register the /voteythumbs command
driftwood.register_application_command({
    name = "voteythumbs",
    description = "Respectfully debate things that probably don't matter",
    options = {
        {
            name = "question",
            description = "The topic of debate",
            type = driftwood.option_string,
            required = true,
        },
    },
    handler = function(interaction)
      
        local question = interaction.options.question
        local message_id = driftwood.message.add(interaction.channel_id, question)
        if not message_id then
            interaction:reply("Failed to send message.", { ephemeral = true })
            return
        end
      
        driftwood.reaction.add(message_id, interaction.channel_id, "👍")
        driftwood.reaction.add(message_id, interaction.channel_id, "👎")

        interaction:reply("Your vote has been setup!", { ephemeral = true })
      
    end,
})
