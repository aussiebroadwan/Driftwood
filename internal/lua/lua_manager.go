package lua

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"log/slog"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"

	"driftwood/internal/lua/bindings"
)

// LuaManager handles loading and executing Lua scripts and binding them to Discord commands/events.
type LuaManager struct {
	LuaState *lua.LState // The Lua VM state
	Bindings []bindings.LuaBinding
}

// NewManager creates a new LuaManager with the given session and Guild ID.
func NewManager(session *discordgo.Session, guildID string) *LuaManager {
	manager := &LuaManager{
		LuaState: lua.NewState(),
		Bindings: []bindings.LuaBinding{
			bindings.NewApplicationCommandBinding(session, guildID),
			// Add more bindings here as needed.
		},
	}

	// register the bindings
	manager.RegisterDiscordModule()
	return manager
}

// LoadScripts loads all Lua scripts from the directory specified in the `LUA_SCRIPTS_PATH` environment variable.
func (m *LuaManager) LoadScripts(path string) error {
	if path == "" {
		return errors.New("LUA_SCRIPTS_PATH is not set")
	}

	slog.Info("Loading Lua scripts", "path", path)

	// Convert the path to an absolute path.
	absPath, err := filepath.Abs(path)
	if err != nil {
		absPath = path // Fallback to relative path if absolute conversion fails.
	}

	// Update `package.path` to include the new path.
	packagePath := m.LuaState.GetField(m.LuaState.GetGlobal("package"), "path").String()
	newPath := filepath.Join(absPath, "?.lua")
	m.LuaState.SetField(m.LuaState.GetGlobal("package"), "path", lua.LString(packagePath+";"+newPath))

	// Walk through the directory and load each Lua script
	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error walking through Lua scripts: %w", err)
		}

		// Skip directories that are not the entry point (init.lua).
		if info.IsDir() {
			initFilePath := filepath.Join(path, "init.lua")
			if _, err := os.Stat(initFilePath); err == nil {
				slog.Info("Loading Lua module", "path", initFilePath)
				if loadErr := m.LuaState.DoFile(initFilePath); loadErr != nil {
					slog.Error("Failed to load Lua module", "path", initFilePath, "error", loadErr)
					return loadErr
				}
			}
			return nil
		}

		// Load single-file commands.
		if filepath.Ext(path) == ".lua" && info.Name() != "init.lua" {
			slog.Info("Loading Lua script", "path", path)
			if loadErr := m.LuaState.DoFile(path); loadErr != nil {
				slog.Error("Failed to load Lua script", "path", path, "error", loadErr)
				return loadErr
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	slog.Info("Lua scripts loaded successfully")
	return nil
}

// DiscordOptionTypes maps human-readable constants to Discord's option type values.
var DiscordOptionTypes = map[string]int{
	"option_subcommand":       1,
	"option_subcommand_group": 2,
	"option_string":           3,
	"option_integer":          4,
	"option_boolean":          5,
	"option_user":             6,
	"option_channel":          7,
	"option_role":             8,
	"option_mentionable":      9,
	"option_number":           10,
	"option_attachment":       11,
}

// RegisterDiscordModule creates a custom loader for `require("discord")`
// and injects the actual Go bindings into the Lua state.
func (m *LuaManager) RegisterDiscordModule() {
	// Loader function for `require("discord")`.
	discordLoader := func(L *lua.LState) int {
		module := L.NewTable()

		// Add constants to the module.
		for key, value := range DiscordOptionTypes {
			module.RawSetString(key, lua.LNumber(value))
		}

		// Register the function bindings.
		for _, binding := range m.Bindings {
			fn := binding.Register(m.LuaState)
			L.SetField(module, binding.Name(), fn)
			slog.Info("Registered binding", "name", binding.Name())
		}

		L.Push(module)
		return 1
	}

	// Register the loader.
	m.LuaState.PreloadModule("discord", discordLoader)
}

func (m *LuaManager) HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	slog.Info("Handling command", "interaction", i.ID)

	// Route the command to the ApplicationCommandBinding.
	for idx := range m.Bindings {
		if err := m.Bindings[idx].HandleCommand(m.LuaState, i); err == nil {
			return // Command was handled successfully
		}
	}

	slog.Warn("Command binding not found",
		"command", i.ApplicationCommandData().Name,
		slog.Any("options", i.ApplicationCommandData().Options),
		slog.Any("bindings", m.Bindings))
}
