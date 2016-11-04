package json

import (
	
	"testing"
	"../fs"
)

func Test_string2Json(t *testing.T) {
	j, err := String2Json(`{"a":1,"b":"abc"}`)
	if err != nil {
		t.Error(err)
		return
	}
	i, err := j.Get("a").Int()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("result:", i)
}