package main

const example = `(Origin == "MOW" || Country == "RU") && (Value >= 100 || Adults == 1)`

const js_loop = `function sum(n) {
	var ret=0;
	var i=0;
	for (i=0; i<=n;i++){
		ret+=i;
	}
	return ret;
}
sum(1000);`

const js_fib = `function fib(n) {
	if (n < 2) return n;
	return fib(n - 2) + fib(n - 1);
}
fib(15);`

func createParams() map[string]interface{} {
	params := make(map[string]interface{})
	params["Origin"] = "MOW"
	params["Country"] = "RU"
	params["Adults"] = 1
	params["Value"] = 100
	return params
}

type Params struct {
	Origin  string
	Country string
	Value   int
	Adults  int
}
