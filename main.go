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
	win2sizeH := L.GetField(win2size, "h")
	fmt.Println("window2.size.h =", win2sizeH)

	// run lua code from go
	if result, err := DoFuncLuaRet("concatL", lua.LString("Go"), lua.LString("Lua")); err != nil {
		fmt.Println("concat error")
	} else {
		fmt.Println("concat = ", result)
	}

	// again run lua code from go
	if _, err := DoFuncLuaRet("printMessageLua", lua.LString("Go")); err != nil {
		fmt.Println("printMessageLua error")
	}

	if err := registerF("_square", square); err != nil {
		fmt.Println("_square error")
	}

	if err := registerF("_summa", summa); err != nil {
		fmt.Println("_summa error")
	}

	//ret 1 params
	if result, err := DoFuncLuaRet("squareNumberL", lua.LNumber(6)); err != nil {
		fmt.Println("squareNumbers error")
	} else {
		fmt.Println("squareNumbers =", result)
	}

	//run 2 params
	if result, err := DoFuncLuaRets("sumNumbersL", 2, lua.LNumber(2), lua.LNumber(6)); err != nil {
		fmt.Println("sumNumbers error")
	} else {
		fmt.Println("sumNumbers1 =", result[0])
		fmt.Println("sumNumbers2 =", result[1])
	}

	// http works
	if result, err := DoFuncLuaRets("getpageL", 2, lua.LString("http://vsi.org.ua")); err != nil {
		fmt.Println("getpage error")
	} else {
		fmt.Println("http code : ", result[0])
		fmt.Println("http size : ", result[1])
	}

}
