# GoLuaTest
Golang, Lua

see comments in main.go

call lua functions uses modified function with slice args
<p>
<code>
func (ls *LState) CallByParams(cp P, args []LValue) error {
	ls.Push(cp.Fn)
	for _, arg := range args {
		ls.Push(arg)
	}

	if cp.Protect {
		return ls.PCall(len(args), cp.NRet, cp.Handler)
	}
	ls.Call(len(args), cp.NRet)
	return nil
}
</code>
</p>