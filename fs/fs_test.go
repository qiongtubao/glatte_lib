package fs

import (
	
	"testing"
)

func Test_readFile(t *testing.T) {
	//for i := 0; i < b.N; i++ {

	err := WriteFile("./test.txt", "aaaa")
	if err != nil {
		t.Error("write error", err)
		return
	}

	bs, err := ReadFile("./test.txt")
	if err != nil {
		t.Error("read error", err)
		return
	}
	t.Log("read:",string(bs))
	//}
}