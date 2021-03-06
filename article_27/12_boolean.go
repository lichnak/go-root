// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//
// Demonstrační příklad číslo 12:
//    Pravdivostní hodnoty.

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "compare.lua"

func compareNumbers(a int, b int) {
	luaVM := lua.NewState()
	log.Println("Lua VM has been initialized")

	defer func() {
		luaVM.Close()
		log.Println("Lua VM has been closed")
	}()

	err := luaVM.DoFile(LuaSource)
	if err != nil {
		log.Fatal(err)
	}

	err = luaVM.CallByParam(lua.P{
		Fn:      luaVM.GetGlobal("compare"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(a), lua.LNumber(b))

	if err != nil {
		log.Fatal(err)
	}

	ret := luaVM.Get(-1)
	luaVM.Pop(1)
	println("Type", ret.Type())
	println("Value", ret.String())
	println("Result", ret == lua.LTrue)
}

func main() {
	compareNumbers(1, 2)
	compareNumbers(1, 1)
}
