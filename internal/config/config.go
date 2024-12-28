package config

import (
	"fmt"
	"os"
	"strconv"

	"log/slog"

	"github.com/joho/godotenv"
)

// Config represents the configuration for the Driftwood bot.
type Config struct {
	DiscordToken   string // Discord bot token
	LuaScriptsPath string // Path to the Lua scripts directory
	GuildID        string // Guild ID (Server ID) for bot commands
}

// Load loads the configuration from environment variables and `.env` files.
// It validates required fields and provides defaults for optional fields.
func Load() (*Config, error) {
	// Load environment variables from a .env file if it exists
	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		slog.Warn("Failed to load .env file", "error", err)
	}

	// Create the config struct
	cfg := &Config{
		DiscordToken:   os.Getenv("DISCORD_TOKEN"),
		LuaScriptsPath: getEnvOrDefault("LUA_SCRIPTS_PATH", "/lua"),
		GuildID:        os.Getenv("GUILD_ID"),
	}

	// Validate required fields
	if err := cfg.validate(); err != nil {
		return nil, err
	}

	slog.Info("Configuration loaded successfully", "LuaScriptsPath", cfg.LuaScriptsPath, "GuildID", cfg.GuildID)
	return cfg, nil
}

// validate ensures that all required configuration fields are set and valid.
func (cfg *Config) validate() error {
	if cfg.DiscordToken == "" {
		return fmt.Errorf("DISCORD_TOKEN is required but not set")
	}
	if cfg.GuildID == "" {
		return fmt.Errorf("GUILD_ID is required but not set")
	}
	if _, err := strconv.ParseUint(cfg.GuildID, 10, 64); err != nil {
		return fmt.Errorf("GUILD_ID must be a valid non-zero integer: %s", cfg.GuildID)
	}
	if _, err := os.Stat(cfg.LuaScriptsPath); os.IsNotExist(err) {
		return fmt.Errorf("LUA_SCRIPTS_PATH does not exist: %s", cfg.LuaScriptsPath)
	}
	return nil
}

// getEnvOrDefault retrieves an environment variable or returns a default value if unset.
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
