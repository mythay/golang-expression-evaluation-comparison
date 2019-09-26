package main

import (
	"testing"

	lua "github.com/yuin/gopher-lua"
)

func Benchmark_lua(b *testing.B) {

	b.ResetTimer()

	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("Origin", lua.LString("MOW"))
	L.SetGlobal("Country", lua.LString("RU"))
	L.SetGlobal("Adults", lua.LNumber(1))
	L.SetGlobal("Value", lua.LNumber(100))
	fn, err := L.LoadString(`return ((Origin == "MOW" or Country == "RU") and (Value >= 100 or Adults == 1))`)
	if err != nil {
		panic(err)
	}
	// L.DoString(`return ((Origin == "MOW" or Country == "RU") and (Value >= 100 or Adults == 1))`)
	for n := 0; n < b.N; n++ {
		L.Push(fn)
		err = L.PCall(0, lua.MultRet, nil)
		if err != nil {
			panic(err)
		}
		L.Pop(1)
	}
	b.StopTimer()
}

func Benchmark_lua_vm(b *testing.B) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		L := lua.NewState()
		defer L.Close()

		L.SetGlobal("Origin", lua.LString("MOW"))
		L.SetGlobal("Country", lua.LString("RU"))
		L.SetGlobal("Adults", lua.LNumber(1))
		L.SetGlobal("Value", lua.LNumber(100))
		fn, err := L.LoadString(`return ((Origin == "MOW" or Country == "RU") and (Value >= 100 or Adults == 1))`)
		if err != nil {
			panic(err)
		}
		// L.DoString(`return ((Origin == "MOW" or Country == "RU") and (Value >= 100 or Adults == 1))`)
		L.Push(fn)
		err = L.PCall(0, lua.MultRet, nil)
		if err != nil {
			panic(err)
		}
		L.Pop(1)
	}
	b.StopTimer()
}

func Benchmark_lua_vm_without_lib(b *testing.B) {

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		L := lua.NewState(lua.Options{SkipOpenLibs: true})
		defer L.Close()

		L.SetGlobal("Origin", lua.LString("MOW"))
		L.SetGlobal("Country", lua.LString("RU"))
		L.SetGlobal("Adults", lua.LNumber(1))
		L.SetGlobal("Value", lua.LNumber(100))
		fn, err := L.LoadString(`return ((Origin == "MOW" or Country == "RU") and (Value >= 100 or Adults == 1))`)
		if err != nil {
			panic(err)
		}
		// L.DoString(`return ((Origin == "MOW" or Country == "RU") and (Value >= 100 or Adults == 1))`)
		L.Push(fn)
		err = L.PCall(0, lua.MultRet, nil)
		if err != nil {
			panic(err)
		}
		L.Pop(1)
	}
	b.StopTimer()
}

func Benchmark_lua_number_fib(b *testing.B) {

	L := lua.NewState()
	defer L.Close()

	fn, err := L.LoadString(`local function fib(n)
    if n < 2 then return n end
    return fib(n - 2) + fib(n - 1)
end
return fib(15)`)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	// L.DoString(`return ((Origin == "MOW" or Country == "RU") and (Value >= 100 or Adults == 1))`)
	for n := 0; n < b.N; n++ {
		L.Push(fn)
		err = L.PCall(0, lua.MultRet, nil)
		if err != nil {
			panic(err)
		}
		val := L.Get(-1)
		L.Pop(1)
		if n, ok := val.(lua.LNumber); ok {
			if n != 610 {
				b.Fatal(n)
			}
		}
	}
	b.StopTimer()
}

func Benchmark_lua_number_loop(b *testing.B) {

	L := lua.NewState()
	defer L.Close()

	fn, err := L.LoadString(`local function sum(n)
	local ret=0
	for i=0,n do
		ret=ret+i
	end
	return ret
end
return sum(1000)`)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		L.Push(fn)
		err = L.PCall(0, lua.MultRet, nil)
		if err != nil {
			panic(err)
		}
		val := L.Get(-1)
		L.Pop(1)
		if n, ok := val.(lua.LNumber); ok {
			if n != 500500 {
				b.Fatal(n)
			}
		}

	}
	b.StopTimer()

}

func Benchmark_lua_number_loop_big(b *testing.B) {

	L := lua.NewState()
	defer L.Close()

	fn, err := L.LoadString(`local function sum(n)
	local ret=0
	for i=0,n do
		ret=ret+i
	end
	return ret
end
return sum(10000)`)
	if err != nil {
		panic(err)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		L.Push(fn)
		err = L.PCall(0, lua.MultRet, nil)
		if err != nil {
			panic(err)
		}
		val := L.Get(-1)
		L.Pop(1)
		if n, ok := val.(lua.LNumber); ok {
			if n != 50005000 {
				b.Fatal(n)
			}
		}

	}
	b.StopTimer()

}
