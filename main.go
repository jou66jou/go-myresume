//隨手測試用
package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Name string `json:"name"`
}

type B struct {
	Age int `json:"age"`
}

type hub struct {
	C chan interface{}
}

func do(c chan interface{}) {
	for k := range c {
		json, _ := json.Marshal(k)
		fmt.Println(string(json))
	}
}

func main() {
	hub := &hub{make(chan interface{})}
	go func() {
		hub.C <- A{"jaseHuang"}
		hub.C <- B{25}
	}()

	go do(hub.C)
	for {
	}
}
