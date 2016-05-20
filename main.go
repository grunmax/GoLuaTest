// LuaTest1 project main.go
package main

import (
	"fmt"
	"net/http"

	"github.com/cjoudrey/gluahttp"
	"github.com/yuin/gopher-lua"
)

var L *lua.LState

func init() {
	L = lua.NewState()
	L.PreloadModule("http", gluahttp.NewHttpModule(&http.Client{}).Loader)
}

func main() {
	defer L.Close()
	if err := L.DoFile("test.lua"); err != nil {
		panic(err)
	}

	// read/write globals vars
	if answer, err := GetIntLua("answer"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer)
	}
	if testString, err := GetStringLua("testString"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(testString)
	}

	L.SetGlobal("receivedString2", lua.LString("Строка из Go"))
	testString := L.GetGlobal("receivedString2")
	fmt.Println(testString)

	// read/write tables fields
	win1title := L.GetField(L.GetGlobal("window"), "title")
	fmt.Println(win1title)

	win2size := L.GetField(L.GetGlobal("window2"), "size")
	win2sizeW := L.GetField(win2size, "w")
	fmt.Println(win2sizeW)
	L.SetField(win2size, "w", lua.LNumber(33))
	win2sizeW = L.GetField(win2size, "w")
	fmt.Println(win2sizeW)

	// run lua code from go
	if result, err := DoFuncLuaRet("concat", []lua.LValue{lua.LString("Go"), lua.LString("Lua")}); err != nil {
		fmt.Println("concat error")
	} else {
		fmt.Println(result)
	}
	if _, err := DoFuncLuaRet("printMessageLua", []lua.LValue{lua.LString("Go")}); err != nil {
		fmt.Println("printMessageLua error")
	}

	// run go code from lua
	L.SetGlobal("squareGO", L.NewFunction(square))
	if result, err := DoFuncLuaRet("squareGO", []lua.LValue{lua.LNumber(5)}); err != nil {
		fmt.Println("squareGO error")
	} else {
		fmt.Println(result)
	}

	//run 2 params
	if result, err := DoFuncLuaRets("sumNumbers", []lua.LValue{lua.LNumber(2), lua.LNumber(6)}, 2); err != nil {
		fmt.Println("sumNumbers error")
	} else {
		fmt.Println(result)
	}

	// http works
	if result, err := DoFuncLuaRet("getpage", []lua.LValue{lua.LString("http://vsi.org.ua")}); err != nil {
		fmt.Println("getpage error")
	} else {
		fmt.Println(result)
	}

}
