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

func summa(L *lua.LState) int { //*
	a := L.ToInt(1) // get first (1) function argument and convert to int
	b := L.ToInt(2)
	ln := lua.LNumber(a + b) // make calculation and cast to LNumber
	L.Push(ln)               // Push it to the stack
	return 1                 // Notify that we pushed one value to the stack
}
