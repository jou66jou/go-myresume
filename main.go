package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	Peers := []string{"123", "456", "abc"}
	addr := []interface{}{}
	b, e := json.Marshal(Peers)
	if e != nil {
		fmt.Println(e)
	}
	json.Unmarshal(b, &addr)
	// fmt.Printf("%+v\n", addr)

	for k, v := range addr {
		fmt.Println(k, v)
	}
}
