package main

import (
	"github.com/yuin/gopher-lua"
)

func square(L *lua.LState) int { //*
	i := L.ToInt(1)          // get first (1) function argument and convert to int
	ln := lua.LNumber(i * i) // make calculation and cast to LNumber
	L.Push(ln)               // Push it to the stack
	return 1                 // Notify that we pushed one value to the stack
}

func DoFuncLuaRet(name string, args []lua.LValue) (lua.LValue, error) {
	luaP := lua.P{
		Fn:      L.GetGlobal(name), // name of Lua function
		NRet:    1,                 // number of returned values
		Protect: true,              // return err or panic
	}
	if err := L.CallByParams(luaP, args); err != nil {
		return lua.LNil, err
	}
	return L.Get(-1), nil
}

func DoFuncLuaRets(name string, args []lua.LValue, returnedcount int) ([]lua.LValue, error) {
	luaP := lua.P{
		Fn:      L.GetGlobal(name), // name of Lua function
		NRet:    returnedcount,     // number of returned values
		Protect: true,              // return err or panic
	}
	if err := L.CallByParams(luaP, args); err != nil {
		return []lua.LValue{}, err
	}
	var result []lua.LValue
	for i := -1 * returnedcount; i < 0; i++ {
		result = append(result, L.Get(i))
	}
	return result, nil
}

func GetIntLua(name string) (int, error) {
	value := L.GetGlobal(name)
	if value == lua.LNil {
		return 0, MyErr{"not found"}
	} else {
		return int(value.(lua.LNumber)), nil
	}
}

func GetStringLua(name string) (string, error) {
	value := L.GetGlobal(name)
	if value == lua.LNil {
		return "", MyErr{"not found"}
	} else {
		return string(value.(lua.LString)), nil
	}
}
