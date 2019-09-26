package main

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func Benchmark_otto(b *testing.B) {
	params := createParams()

	vm := otto.New()

	script, err := vm.Compile("", example)
	if err != nil {
		b.Fatal(err)
	}

	_ = vm.Set("Origin", params["Origin"])
	_ = vm.Set("Country", params["Country"])
	_ = vm.Set("Adults", params["Adults"])
	_ = vm.Set("Value", params["Value"])

	var out otto.Value

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(script)
	}
	b.StopTimer()

	if err != nil {
		b.Fatal(err)
	}
	ok, err := out.ToBoolean()
	if err != nil {
		b.Fatal(err)
	}
	if !ok {
		b.Fail()
	}
}

func Benchmark_otto_vm_create(b *testing.B) {
	params := createParams()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		vm := otto.New()

		script, err := vm.Compile("", example)
		if err != nil {
			b.Fatal(err)
		}

		_ = vm.Set("Origin", params["Origin"])
		_ = vm.Set("Country", params["Country"])
		_ = vm.Set("Adults", params["Adults"])
		_ = vm.Set("Value", params["Value"])

		var out otto.Value

		out, err = vm.Run(script)
		if err != nil {
			b.Fatal(err)
		}
		ok, err := out.ToBoolean()
		if err != nil {
			b.Fatal(err)
		}
		if !ok {
			b.Fail()
		}
	}
	b.StopTimer()

}

func Benchmark_otto_number_fib(b *testing.B) {
	vm := otto.New()

	script, err := vm.Compile("", js_fib)
	if err != nil {
		b.Fatal(err)
	}

	var out otto.Value

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(script)
		if err != nil {
			b.Fatal(err)
		}
		n, err := out.ToInteger()
		if err != nil {
			b.Fatal(err)
		}
		if n != 610 {
			b.Fatal(n)
		}
	}
	b.StopTimer()

}

func Benchmark_otto_number_loop(b *testing.B) {
	vm := otto.New()

	script, err := vm.Compile("", js_loop)
	if err != nil {
		b.Fatal(err)
	}

	var out otto.Value

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		out, err = vm.Run(script)
		if err != nil {
			b.Fatal(err)
		}
	}
	b.StopTimer()

	n, err := out.ToInteger()
	if err != nil {
		b.Fatal(err)
	}
	if n != 500500 {
		b.Fatal(n)
	}

}
