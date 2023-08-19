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
	if err := L.DoFile("Lua/test.lua"); err != nil {
		panic(err)
	}

	// read globals vars
	if answer, err := GetIntLua("answer"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("answer =", answer)
	}
	if testString, err := GetStringLua("testString"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("testString =", testString)
	}

	// write global var
	L.SetGlobal("receivedString", lua.LString("Строка из Go"))
	if testString, err := GetStringLua("receivedString"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("receivedString = ", testString)
	}

	//create global var
	L.SetGlobal("receivedString2", lua.LString("Ещё строка из Go"))
	testString := L.GetGlobal("receivedString2") // with no check
	fmt.Println("new receivedString2 = ", testString)

	// read/write tables fields
	win1title := L.GetField(L.GetGlobal("window"), "title")
	fmt.Println("window.title =", win1title)

	win2size := L.GetField(L.GetGlobal("window2"), "size")
	win2sizeW := L.GetField(win2size, "w")
	fmt.Println("window2.size.w =", win2sizeW)
	L.SetField(win2size, "w", lua.LNumber(33))
	win2sizeW = L.GetField(win2size, "w")
	fmt.Println("window2.size.w =", win2sizeW)

	// run lua code from go
	if result, err := DoFuncLuaRet("concat", lua.LString("Go"), lua.LString("Lua")); err != nil {
		fmt.Println("concat error")
	} else {
		fmt.Println("concat = ", result)
	}

	// again run lua code from go
	if _, err := DoFuncLuaRet("printMessageLua", lua.LString("Go")); err != nil {
		fmt.Println("printMessageLua error")
	}

	// set go code to lua
	L.SetGlobal("squareGO", L.NewFunction(square)) // square is go fun
	// check setted go code
	if result, err := DoFuncLuaRet("squareGO", lua.LNumber(5)); err != nil {
		fmt.Println("squareGO error")
	} else {
		fmt.Println("squareGO fun: ", result)
	}

	//run 2 params
	if result, err := DoFuncLuaRets("sumNumbers", 2, lua.LNumber(2), lua.LNumber(6)); err != nil {
		fmt.Println("sumNumbers error")
	} else {
		fmt.Println("sumNumbers =", result)
	}

	// http works
	if result, err := DoFuncLuaRet("getpage", lua.LString("http://vsi.org.ua")); err != nil {
		fmt.Println("getpage error")
	} else {
		fmt.Println("http fun : ", result)
	}

}
