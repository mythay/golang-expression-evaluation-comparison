package main

import (
	"testing"

	"github.com/d5/tengo/script"
)

// func Benchmark_tengo(b *testing.B) {
// 	params := createParams()

// 	vm := goja.New()
// 	program, err := goja.Compile("", example, false)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	vm.Set("Origin", params["Origin"])
// 	vm.Set("Country", params["Country"])
// 	vm.Set("Adults", params["Adults"])
// 	vm.Set("Value", params["Value"])

// 	var out goja.Value

// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		out, err = vm.RunProgram(program)
// 	}
// 	b.StopTimer()

// 	if err != nil {
// 		b.Fatal(err)
// 	}
// 	if !out.ToBoolean() {
// 		b.Fail()
// 	}
// }

// func Benchmark_tengo_vm_create(b *testing.B) {
// 	params := createParams()
// 	b.ResetTimer()

// 	for n := 0; n < b.N; n++ {
// 		vm := goja.New()
// 		program, err := goja.Compile("", example, false)
// 		if err != nil {
// 			b.Fatal(err)
// 		}

// 		vm.Set("Origin", params["Origin"])
// 		vm.Set("Country", params["Country"])
// 		vm.Set("Adults", params["Adults"])
// 		vm.Set("Value", params["Value"])

// 		var out goja.Value
// 		out, err = vm.RunProgram(program)
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 		if !out.ToBoolean() {
// 			b.Fail()
// 		}
// 	}
// 	b.StopTimer()

// }

// func Benchmark_goja_number_fib(b *testing.B) {

// 	vm := goja.New()
// 	program, err := goja.Compile("", js_fib, false)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	var out goja.Value

// 	b.ResetTimer()
// 	for n := 0; n < b.N; n++ {
// 		out, err = vm.RunProgram(program)
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 		if out.ToInteger() != 55 {
// 			b.Fatal(out)
// 		}
// 	}
// 	b.StopTimer()

// }

func Benchmark_tengo_number_loop(b *testing.B) {

	vm := script.New([]byte(`
	ret:=0
	sum:=func(n){
		for i:=0; i<n; i++{
			ret+=i
		}
		return ret
	}
	sum(1000)
	`))
	program, err := vm.Compile()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = program.Run()
		if err != nil {
			b.Fatal(err)
		}
		// if out.ToInteger() != 500500 {
		// 	b.Fatal(out)
		// }
	}
	b.StopTimer()

}
