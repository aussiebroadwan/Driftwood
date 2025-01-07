package bot

import (
	"log/slog"

	"driftwood/internal/lua"

	"github.com/bwmarrin/discordgo"
)

// Bot represents the Discord bot instance.
type Bot struct {
	Session *discordgo.Session // Discord session
	GuildID string             // Guild ID (Server ID) for command registration

	luaMgr *lua.LuaManager // Lua script manager
}

// NewBot initializes a new bot instance with the given Discord token.
// Returns a Bot instance or an error if initialization fails.
func NewBot(token string) (*Bot, error) {
	slog.Info("Creating a new bot session")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		slog.Error("Failed to create Discord session", "error", err)
		return nil, err
	}
	return &Bot{Session: session}, nil
}

// SetGuildID sets the Guild ID (Server ID) for command registration.
func (b *Bot) SetGuildID(guildID string) {
	b.GuildID = guildID
}

// Start opens the Discord WebSocket connection and registers event handlers.
// It also loads Lua scripts to initialize commands and events.
func (b *Bot) Start(path string) error {
	slog.Info("Starting bot session")

	// Load Lua scripts and register commands
	if err := b.loadLuaScripts(path); err != nil {
		slog.Error("Failed to load Lua scripts", "error", err)
		return err
	}

	// Register the command interaction handler
	b.Session.AddHandler(b.luaMgr.ReadyHandler)
	b.Session.AddHandler(b.commandHandler)

	// Open the session
	if err := b.Session.Open(); err != nil {
		slog.Error("Failed to open Discord session", "error", err)
		return err
	}

	slog.Info("Bot started successfully")
	return nil
}

// Stop gracefully closes the Discord session.
func (b *Bot) Stop() {
	slog.Info("Stopping bot session")
	err := b.Session.Close()
	if err != nil {
		slog.Error("Failed to close Discord session", "error", err)
	}
}

// commandHandler processes incoming interactions and routes them to Lua-defined commands.
func (b *Bot) commandHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	slog.Info("Received interaction", "type", i.Type, "name", i.Data)

	// Execute the corresponding Lua command
	b.luaMgr.HandleCommand(s, i)
}

// loadLuaScripts loads all Lua scripts, registers commands, and binds events.
func (b *Bot) loadLuaScripts(path string) error {
	// Initialize the Lua manager with the bot's session and Guild ID
	b.luaMgr = lua.NewManager(b.Session, b.GuildID)

	// Load Lua scripts from the configured directory
	if err := b.luaMgr.LoadScripts(path); err != nil {
		return err
	}

	slog.Info("Lua scripts loaded and commands/events registered successfully")
	return nil
}
