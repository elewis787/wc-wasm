package main

import (
	"syscall/js"

	"github.com/elewis787/wc-wasm/wasm/rpc"
)

func new(this js.Value, args []js.Value) interface{} {
	http := rpc.NewHTTP()
	return js.ValueOf(map[string]interface{}{
		"Get": http.Get(),
	})
}

func main() {
	c := make(chan struct{})
	js.Global().Set("http", js.FuncOf(new))
	<-c
}
