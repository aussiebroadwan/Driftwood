package utils

import (
	"sync"

	lua "github.com/yuin/gopher-lua"
)

type luaTask func(L *lua.LState)

type LuaRunner struct {
	L     *lua.LState
	tasks chan luaTask
}

var runner *LuaRunner
var once sync.Once

func GetLuaRunner() *LuaRunner {
	once.Do(func() {
		L := lua.NewState()
		r := &LuaRunner{
			L:     L,
			tasks: make(chan luaTask, 100),
		}
		runner = r

		go runner.loop()
	})
	return runner
}

func (r *LuaRunner) loop() {
	for task := range r.tasks {
		task(r.L)
	}
}

// schedule a call
func (r *LuaRunner) Do(task luaTask) {
	r.tasks <- task
}
