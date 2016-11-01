package fs

import (
	"fmt"
	"testing"
)

func Benchmark_readFile(t *testing.B) {
	//for i := 0; i < b.N; i++ {

	err := WriteFile("./test.txt", "aaaa")
	if err != nil {
		fmt.Println("write error", err)
		return
	}

	bs, err := ReadFile("./test.txt")
	if err != nil {
		fmt.Println("read error", err)
		return
	}
	fmt.Println("read:",string(bs))
	//}
}