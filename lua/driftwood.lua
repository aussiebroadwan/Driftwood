---@meta
--- Driftwood: A modular framework for creating Discord bots using Lua scripts.
--- This file provides type annotations and documentation for all Driftwood bindings.
--- 
--- Visit the Driftwood GitHub repository for more information and updates:
--- https://github.com/lcox74/driftwood

--- Driftwood namespace
local driftwood = {
    state = {},
    timer = {},
    log = {},
    option = {},
    message = {},
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
--- @field reply_with_action fun(self: InteractionBase, content: string, components: InteractionComponents[], options?: InteractionReplyOptions) Replies with action components (e.g., buttons).

--- CommandInteraction class for handling command interactions.
--- Extends the base Interaction class and includes options.
--- @class CommandInteraction : InteractionBase
--- @field options table<string, any> Arguments/options passed to the command interaction.

--- EventInteraction class for handling event interactions (e.g., custom IDs).
--- Extends the base Interaction class and includes data.
--- @class EventInteraction : InteractionBase
--- @field data table<string, string> Parsed regex groups from the custom ID.

--- InteractionReplyOptions class for defining reply options.
--- @class InteractionReplyOptions
--- @field ephemeral? boolean Whether the reply should be ephemeral (default: false).
--- @field mention? boolean Whether to mention the user in the reply (default: true).

--- InteractionComponents class for defining message components (e.g., buttons).
--- @class InteractionComponents
--- @field type string The type of component (e.g., "button").
--- @field label string The label text for the component.
--- @field custom_id string The custom ID for the component.
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
--- @return InteractionComponents button The new button component.
function driftwood.new_button(label, custom_id) end

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
--- @param components? InteractionComponents[] Optional components to include with the message.
--- @return string|nil message_id The ID of the sent message, or nil if failed.
function driftwood.message.add(channel_id, content, components) end

--- Edit an existing message.
--- @param message_id string The ID of the message to edit.
--- @param channel_id string The ID of the channel containing the message.
--- @param content string The new message content.
--- @param components? InteractionComponents[] Optional components to include with the message.
--- @return boolean success Whether the edit was successful.
function driftwood.message.edit(message_id, channel_id, content, components) end

--- Delete a message.
--- @param message_id string The ID of the message to delete.
--- @param channel_id string The ID of the channel containing the message.
--- @return boolean success Whether the deletion was successful.
function driftwood.message.delete(message_id, channel_id) end

--- Command Registration

--- Register an application command.
--- @param command Command A table defining the command, its options, and handlers.
function driftwood.register_application_command(command) end

--- Register an interaction event.
--- @param custom_id string The custom ID or regex for the interaction.
--- @param handler fun(interaction: EventInteraction) The handler function for the interaction.
function driftwood.register_interaction(custom_id, handler) end

--- Discord Option Types

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
