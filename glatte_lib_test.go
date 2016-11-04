package glatte_lib

import (
	"testing"
	
)

func Test_IsArray(t *testing.T) {
	var a = []byte{'a', 'b', 'c'}
	t.Log(IsArray(a))
}

func Test_Type(t *testing.T) {
	var a = []byte{'a', 'b', 'c'}
	t.Log(Type(a))
}