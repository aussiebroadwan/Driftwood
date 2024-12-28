package main

import (
	"driftwood/internal/bot"
	"driftwood/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	// Initialize the bot
	b, err := bot.NewBot(cfg.DiscordToken)
	if err != nil {
		slog.Error("Failed to create bot", "error", err)
		os.Exit(1)
	}

	// Pass GuildID to bot for command registration
	b.SetGuildID(cfg.GuildID)

	// Start the bot
	go func() {
		if err := b.Start(cfg.LuaScriptsPath); err != nil {
			slog.Error("Failed to start bot", "error", err)
			os.Exit(1)
		}
	}()
	slog.Info("Bot is running", "GuildID", cfg.GuildID)

	// Wait for termination signal
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	// Stop the bot gracefully
	slog.Info("Shutting down bot")
	b.Stop()
}
