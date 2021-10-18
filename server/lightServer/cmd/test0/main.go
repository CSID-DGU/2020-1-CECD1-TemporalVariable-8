package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type test struct {

	T time.Time `json:"t,omitempty"`
}
func main() {
	var a = test{
		//T: time.Now(),
	}
	m, _ := json.Marshal(a)
	fmt.Println(string(m))
	fmt.Println(time.Time{}.Location())
}
