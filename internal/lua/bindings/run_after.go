package bindings

import (
	"driftwood/internal/lua/utils"
	"fmt"
	"log/slog"
	"time"

	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

// RunAfterBinding implements the `run_after` Lua function.
type RunAfterBinding struct{}

// NewRunAfterBinding creates a new RunAfterBinding.
func NewRunAfterBinding() *RunAfterBinding {
	slog.Debug("Creating new RunAfterBinding")
	return &RunAfterBinding{}
}

// Name returns the name of the binding for global registration in Lua.
func (b *RunAfterBinding) Name() string {
	return "run_after"
}

func (b *RunAfterBinding) SetSession(session *discordgo.Session) {}

// Register creates the `run_after` Lua function and adds it to the Lua state.
func (b *RunAfterBinding) Register() lua.LGFunction {
	return func(L *lua.LState) int {
		// Validate arguments
		fn := L.CheckFunction(1)         // First argument: Lua function
		delaySeconds := L.CheckNumber(2) // Second argument: Delay in seconds
		if delaySeconds < 0 {
			L.ArgError(2, "delay must be a non-negative number")
			return 0
		}

		// Generate a unique global name for the function
		globalName := fmt.Sprintf("__run_after_%d", time.Now().UnixNano())
		L.SetGlobal(globalName, fn)

		// Start a goroutine to delay and call the function
		go func(globalName string) {
			time.Sleep(time.Duration(float64(delaySeconds) * float64(time.Second)))

			utils.GetLuaRunner().Do(func(L *lua.LState) {

				// Lock the Lua state for execution
				if err := L.CallByParam(lua.P{
					Fn:      L.GetGlobal(globalName), // Retrieve the function by global name
					NRet:    0,                       // No return values
					Protect: true,                    // Catch errors
				}); err != nil {
					slog.Error("Failed to execute delayed Lua function", "error", err)
				}

				// Remove the global function to clean up
				L.SetGlobal(globalName, lua.LNil)
			})
		}(globalName)

		return 0
	}
}

// HandleInteraction is not applicable for this binding.
func (b *RunAfterBinding) HandleInteraction(interaction *discordgo.InteractionCreate) error {
	// This binding does not handle interactions
	return nil
}

func (b *RunAfterBinding) CanHandleInteraction(interaction *discordgo.InteractionCreate) bool {
	return false
}
