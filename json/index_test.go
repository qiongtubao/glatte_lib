package json

import (
	"fmt"
	"testing"
	"../fs"
)

func Benchmark_readFile(b *testing.B) {
	j, err := String2Json(`{"a":1,"b":"abc"}`)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	i, err := j.Get("a").Int()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("result:", i)
}