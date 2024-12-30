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
	"driftwood/internal/lua/utils"
)

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

// LuaManager handles loading and executing Lua scripts and binding them to Discord commands/events.
type LuaManager struct {
	LuaState     *lua.LState // The Lua VM state
	Bindings     map[string][]bindings.LuaBinding
	StateManager *utils.StateManager
}

// NewManager creates a new LuaManager with the given session and Guild ID.
func NewManager(session *discordgo.Session, guildID string) *LuaManager {
	sm := utils.NewStateManager()
	manager := &LuaManager{
		LuaState:     lua.NewState(),
		StateManager: sm,
		Bindings:     make(map[string][]bindings.LuaBinding),
	}

	manager.RegisterBindings(session, guildID)

	// register the bindings
	manager.RegisterDiscordModule()
	return manager
}

// RegisterBindings initializes grouped Lua bindings.
func (m *LuaManager) RegisterBindings(session *discordgo.Session, guildID string) {
	m.Bindings = map[string][]bindings.LuaBinding{
		"default": {
			bindings.NewApplicationCommandBinding(session, guildID),
			bindings.NewInteractionEventBinding(session),
			bindings.NewNewButtonBinding(),
		},
		"timer": {
			bindings.NewRunAfterBinding(),
		},
		"state": {
			bindings.NewStateBindingGet(m.StateManager),
			bindings.NewStateBindingSet(m.StateManager),
			bindings.NewStateBindingClear(m.StateManager),
		},
		"message": {
			bindings.NewMessageBindingAdd(session),
			bindings.NewMessageBindingEdit(session),
			bindings.NewMessageBindingDelete(session),
		},
		"option": {
			bindings.NewNewOptionStringBinding(),
			bindings.NewNewOptionNumberBinding(),
			bindings.NewNewOptionBoolBinding(),
		},
	}

	slog.Info("Lua bindings registered successfully")
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
				slog.Debug("Loading Lua module", "path", initFilePath)
				if loadErr := m.LuaState.DoFile(initFilePath); loadErr != nil {
					slog.Error("Failed to load Lua module", "path", initFilePath, "error", loadErr)
					return loadErr
				}
			}
			return nil
		}

		// Load single-file commands.
		if filepath.Ext(path) == ".lua" && info.Name() != "init.lua" {
			slog.Debug("Loading Lua script", "path", path)
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

// RegisterDiscordModule creates a custom loader for `require("driftwood")`
// and injects the actual Go bindings into the Lua state.
func (m *LuaManager) RegisterDiscordModule() {
	// Loader function for `require("driftwood")`.
	discordLoader := func(L *lua.LState) int {
		module := L.NewTable()

		// Add constants to the module.
		for key, value := range DiscordOptionTypes {
			module.RawSetString(key, lua.LNumber(value))
		}

		// Register the function bindings.
		for groupName, group := range m.Bindings {

			if groupName == "default" {
				for _, binding := range group {
					fn := binding.Register(m.LuaState)
					L.SetField(module, binding.Name(), fn)
					slog.Info("Registered binding", "name", binding.Name())
				}

				continue
			}

			// Create a sub-table for the group.
			subTable := L.NewTable()
			for _, binding := range group {
				fn := binding.Register(m.LuaState)
				L.SetField(subTable, binding.Name(), fn)
				slog.Info("Registered binding", "name", binding.Name())
			}
			L.SetField(module, groupName, subTable)
		}

		addLogging(L, module)

		L.Push(module)
		return 1
	}

	// Register the loader.
	m.LuaState.PreloadModule("driftwood", discordLoader)
}

func addLogging(L *lua.LState, module *lua.LTable) {

	logTable := L.NewTable()

	L.SetField(logTable, "debug", L.NewFunction(func(L *lua.LState) int {
		slog.Debug("Lua debug", "message", L.CheckString(1))
		return 0
	}))
	L.SetField(logTable, "info", L.NewFunction(func(L *lua.LState) int {
		slog.Info("Lua info", "message", L.CheckString(1))
		return 0
	}))
	L.SetField(logTable, "error", L.NewFunction(func(L *lua.LState) int {
		slog.Error("Lua error", "message", L.CheckString(1))
		return 0
	}))

	L.SetField(module, "log", logTable)
}

func (m *LuaManager) HandleCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Route the command to the ApplicationCommandBinding.
	for groupIdx := range m.Bindings {
		for idx := range m.Bindings[groupIdx] {
			if m.Bindings[groupIdx][idx].CanHandleInteraction(i) {
				slog.Debug("Binding matched for interaction", "binding", m.Bindings[groupIdx][idx].Name())
				if err := m.Bindings[groupIdx][idx].HandleInteraction(m.LuaState, i); err == nil {
					return // Command was handled successfully
				} else {
					slog.Warn("Error handling interaction with binding", "binding", m.Bindings[groupIdx][idx].Name(), "error", err)
				}
			}
		}
	}

	slog.Warn("Command binding not found", "interaction_id", i.ID)
}
