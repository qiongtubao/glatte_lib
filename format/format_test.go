package format

import (
	"fmt"
	"testing"
	"../json"
)

func Benchmark_readFile(t *testing.B) {
	//for i := 0; i < b.N; i++ {
	j  := make(map[string]interface{})
	j["a"] = "1"
	j["b"] = "3"
	
	str := templateStringFormatByMap("{{a}} + {{b}}", j)
	fmt.Println(str)
}

func Benchmark_format(t *testing.B) {
	m  := make(map[string]interface{})
	m["a"] = 1
	m["b"] = 3
	j := new(json.JSON)
	j.SetData(m) 
	str := templateStringFormatByJSON("{{a}} + {{b}}", j)
	fmt.Println(str)
}