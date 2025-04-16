package utils

import (
	"github.com/bwmarrin/discordgo"
	lua "github.com/yuin/gopher-lua"
)

func NewOptionRegister(optionType discordgo.ApplicationCommandOptionType) lua.LGFunction {
	return func(L *lua.LState) int {
		argCount := L.GetTop()
		buttonTable := L.NewTable()

		switch argCount {
		case 3:
			buttonTable.RawSetString("name", lua.LString(L.CheckString(1)))
			buttonTable.RawSetString("description", lua.LString(L.CheckString(2)))
			buttonTable.RawSetString("required", lua.LBool(L.CheckBool(3)))
		case 2:
			buttonTable.RawSetString("name", lua.LString(L.CheckString(1)))
			buttonTable.RawSetString("description", lua.LString(L.CheckString(2)))
			buttonTable.RawSetString("required", lua.LFalse)
		default:
			L.ArgError(1, "invalid arguments, expected (name, custom_id [, style])")
		}

		// Create a Table for the button
		buttonTable.RawSetString("type", lua.LNumber(optionType))

		L.Push(buttonTable)
		return 1
	}
}
