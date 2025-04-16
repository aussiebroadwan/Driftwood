---@meta
--- Driftwood: A modular framework for creating Discord bots using Lua scripts.
--- This file provides type annotations and documentation for all Driftwood bindings.
--- 
--- Visit the Driftwood GitHub repository for more information and updates:
--- https://github.com/aussiebroadwan/driftwood

--- Driftwood namespace
local driftwood = {
    state = {},
    timer = {},
    log = {},
    option = {},
    message = {},
    reaction = {},
    channel = {},
}

--- Classes

--- User Interaction class 
--- @class User
--- @field id string The unique ID of the user.
--- @field username string The username of the user.
--- @field global_name string The global name of the user.
--- @field discriminator string The discriminator of the user.
--- @field avatar string The avatar URL of the user.

--- Base Interaction class for handling interactions.
--- @class InteractionBase
--- @field interaction_id string The unique ID of the interaction.
--- @field channel_id string The ID of the channel where the interaction occurred.
--- @field user User The user who triggered the interaction.
--- @field reply fun(self: InteractionBase, content: string, options?: InteractionReplyOptions) Replies to the interaction.

--- CommandInteraction class for handling command interactions.
--- Extends the base Interaction class and includes options.
--- @class CommandInteraction : InteractionBase
--- @field options table<string, any> Arguments/options passed to the command interaction.

--- EventInteraction class for handling event interactions (e.g., custom IDs).
--- Extends the base Interaction class and includes data.
--- @class EventInteraction : InteractionBase
--- @field data table<string, string>|nil Parsed regex groups from the custom ID.
--- @field values string[]|nil The values selected in a select menu.

--- InteractionReplyOptions class for defining reply options.
--- @class InteractionReplyOptions
--- @field ephemeral? boolean Whether the reply should be ephemeral (default: false).
--- @field mention? boolean Whether to mention the user in the reply (default: true).
--- @field components? InteractionComponents[] Optional components to include in the reply.
--- @field embed? MessageEmbed Optional embed to include in the reply.

--- MessageOptions class for defining message options.
--- @class MessageOptions
--- @field components? InteractionComponents[] Optional components to include in the message.
--- @field embed? MessageEmbed Optional embed to include in the message.

--- MessageEmbed class for defining embeds in messages.
--- @class MessageEmbed
--- @field title? string The title of the embed.
--- @field description? string The description text.
--- @field url? string The URL associated with the embed.
--- @field color? number The color code of the embed.
--- @field image? table The embed's image.
--- @field image.url? string The URL of the embed image.
--- @field thumbnail? table The embed's thumbnail.
--- @field thumbnail.url? string The URL of the thumbnail image.
--- @field footer? table The footer information.
--- @field footer.text? string The footer text.
--- @field footer.icon_url? string The URL for the footer icon.
--- @field author? table The author information.
--- @field author.name? string The name of the author.
--- @field author.url? string The URL for the author.
--- @field author.icon_url? string The URL for the author's icon.
--- @field fields? MessageEmbedField[] Array of fields to include in the embed.

--- MessageEmbedField class representing a field within an embed.
--- @class MessageEmbedField
--- @field name string The name of the field.
--- @field value string The value of the field.
--- @field inline? boolean Whether the field should be displayed inline (default: false).

--- InteractionComponents class for defining message components (e.g., buttons).
--- @class InteractionComponents
--- @field type string The type of component (e.g., "button").
--- @field label string The label text for the component.
--- @field custom_id string The custom ID for the component.
--- @field disabled? boolean Optional flag to disable the component.
--- @field style? number Optional style value for buttons.

--- Command class for defining application commands.
--- @class Command
--- @field name string The name of the command.
--- @field description string The description of the command.
--- @field options? CommandOption[] Optional array of options or subcommands.
--- @field handler? fun(interaction: CommandInteraction) Function to handle the command.

--- CommandOption class for defining options within commands.
--- @class CommandOption
--- @field name string The name of the option or subcommand.
--- @field description string The description of the option or subcommand.
--- @field type number The type of the option (see `driftwood.option_*`).
--- @field required? boolean Whether the option is required (default: false).
--- @field options? CommandOption[] Optional sub-options for subcommands.
--- @field handler? fun(interaction: CommandInteraction) Optional handler for subcommands.

--- SelectOption class for defining options within select menus.
--- @class SelectOption
--- @field label string The label of the option.
--- @field value string The value of the option.

--- State Management

--- Set a value in the bot's state with an optional expiry time.
--- @param key string The key to store the value under.
--- @param value any The value to store.
--- @param expiry? number The expiry time in seconds (optional).
function driftwood.state.set(key, value, expiry) end

--- Get a value from the bot's state.
--- @param key string The key to retrieve the value for.
--- @return any|nil value The value stored under the key, or nil if not found.
function driftwood.state.get(key) end

--- Clear a value from the bot's state.
--- @param key string The key to clear.
function driftwood.state.clear(key) end

--- Timer Functions

--- Run a function after a specified number of seconds.
--- @param callback fun() The function to execute.
--- @param seconds number The delay time in seconds.
function driftwood.timer.run_after(callback, seconds) end

--- Logging Functions

--- Log a debug-level message.
--- @param message string The message to log.
function driftwood.log.debug(message) end

--- Log an info-level message.
--- @param message string The message to log.
function driftwood.log.info(message) end

--- Log an error-level message.
--- @param message string The message to log.
function driftwood.log.error(message) end

--- Discord Interaction Action Row

--- Create a new instance of a button in an action row.
--- @param label string The label text for the button.
--- @param custom_id string The custom ID for the button.
--- @param disabled? boolean Optional flag to disable the button.
--- @return InteractionComponents button The new button component.
function driftwood.new_button(label, custom_id, disabled) end

--- Create a new instance of a select menu option.
--- @param label string The label of the option.
--- @param value string The custom ID for the option.
--- @return SelectOption opt The new select menu option.
function driftwood.new_selectmenu_opt(label, value) end

--- Create a new instance of a select menu in an action row.
--- @param placeholder string The placeholder text for the select menu.
--- @param custom_id string The custom ID for the select menu.
--- @param options SelectOption[] The options for the select menu.
--- @param disabled boolean Whether the select menu is disabled.
--- @return InteractionComponents selectmenu The new select menu component.
function driftwood.new_selectmenu(placeholder, custom_id, options, disabled) end

--- Options Functions

--- Create a new string option for a command.
--- @param label string The label of the option.
--- @param description string The description of the option.
--- @param required? boolean Whether the option is required (default: false).
--- @return CommandOption option The new string option.
function driftwood.option.new_string(label, description, required) end

--- Create a new boolean option for a command.
--- @param label string The label of the option.
--- @param description string The description of the option.
--- @param required? boolean Whether the option is required (default: false).
--- @return CommandOption option The new boolean option.
function driftwood.option.new_bool(label, description, required) end

--- Create a new number option for a command.
--- @param label string The label of the option.
--- @param description string The description of the option.
--- @param required? boolean Whether the option is required (default: false).
--- @return CommandOption option The new number option.
function driftwood.option.new_number(label, description, required) end

--- Message Functions

--- Add a message to a channel.
--- @param channel_id string The ID of the channel to send the message to.
--- @param content string The message content.
--- @param options? MessageOptions Optional options for the message.
--- @return string|nil message_id The ID of the sent message, or nil if failed.
function driftwood.message.add(channel_id, content, options) end

--- Edit an existing message.
--- @param message_id string The ID of the message to edit.
--- @param channel_id string The ID of the channel containing the message.
--- @param content string The new message content.
--- @param options? MessageOptions Optional options for the message.
--- @return boolean success Whether the edit was successful.
function driftwood.message.edit(message_id, channel_id, content, options) end

--- Delete a message.
--- @param message_id string The ID of the message to delete.
--- @param channel_id string The ID of the channel containing the message.
--- @return boolean success Whether the deletion was successful.
function driftwood.message.delete(message_id, channel_id) end

--- Reaction Functions
--- These functions provide support for adding and removing reactions on messages.

--- Add a reaction to a message.
--- @param message_id string The ID of the message to react to.
--- @param channel_id string The ID of the channel where the message is located.
--- @param reaction_emoji string The emoji to add as a reaction.
--- @return boolean success Whether the reaction was successfully added.
function driftwood.reaction.add(message_id, channel_id, reaction_emoji) end

--- Remove a reaction from a message.
--- @param message_id string The ID of the message from which to remove the reaction.
--- @param channel_id string The ID of the channel where the message is located.
--- @param reaction_emoji string The emoji to remove as a reaction.
--- @return boolean success Whether the reaction was successfully removed.
function driftwood.reaction.remove(message_id, channel_id, reaction_emoji) end


--- Channel Functions

--- Get a channel by name.
--- @param channel_name string The name of the channel.
--- @return string|nil channel_id The ID of the channel, or nil if not found.
function driftwood.channel.get(channel_name) end

--- Command Registration

--- Register an application command.
--- @param command Command A table defining the command, its options, and handlers.
function driftwood.register_application_command(command) end

--- Register an interaction event.
--- @param custom_id string The custom ID or regex for the interaction.
--- @param handler fun(interaction: EventInteraction) The handler function for the interaction.
function driftwood.register_interaction(custom_id, handler) end


--- Register an On Ready event handler.
--- @param handler fun() The handler function for the interaction.
function driftwood.on_ready(handler) end

--- Enum for Discord application command option types.
--- @enum
driftwood.option_subcommand = 1
driftwood.option_subcommand_group = 2
driftwood.option_string = 3
driftwood.option_integer = 4
driftwood.option_boolean = 5
driftwood.option_user = 6
driftwood.option_channel = 7
driftwood.option_role = 8
driftwood.option_mentionable = 9
driftwood.option_number = 10
driftwood.option_attachment = 11

return driftwood
