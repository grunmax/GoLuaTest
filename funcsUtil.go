package main

import (
	"github.com/yuin/gopher-lua"
)

func registerF(name string, fn lua.LGFunction) error {
	L.SetGlobal(name, L.NewFunction(fn)) // square is go fun
	if _, err := DoFuncLuaRet(name, lua.LNumber(5)); err != nil {
		return err
	} else {
		return nil
	}
}

func DoFuncLuaRet(name string, args ...lua.LValue) (lua.LValue, error) {
	luaP := lua.P{
		Fn:      L.GetGlobal(name), // name of Lua function
		NRet:    1,                 // number of returned values
		Protect: true,              // return err or panic
	}
	if err := L.CallByParam(luaP, args...); err != nil {
		return lua.LNil, err
	}
	return L.Get(-1), nil
}

func DoFuncLuaRets(name string, returnedcount int, args ...lua.LValue) ([]lua.LValue, error) {
	luaP := lua.P{
		Fn:      L.GetGlobal(name), // name of Lua function
		NRet:    returnedcount,     // number of returned values
		Protect: true,              // return err or panic
	}
	if err := L.CallByParam(luaP, args...); err != nil {
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
