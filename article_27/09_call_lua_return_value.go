// Seriál "Programovací jazyk Go"
//
// Dvacátá sedmá část
//
// Demonstrační příklad číslo 9:
//    Zavolání Lua funkce z Go s předáním parametrů, přečtení výsledku.
//

package main

import (
	"github.com/yuin/gopher-lua"
	"log"
)

const LuaSource = "add.lua"

func main() {
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
		Fn:      luaVM.GetGlobal("add"),
		NRet:    1,
		Protect: true,
	}, lua.LNumber(1), lua.LNumber(2))

	if err != nil {
		log.Fatal(err)
	}

	ret := luaVM.Get(-1)
	luaVM.Pop(1)
	println("Type", ret.Type())
	println("Value", ret.String())
}
