package helloworld

import (
	"testing"
)

func Test_hello_one(t *testing.T) {
	if "hello one" == Hello("one") {
		t.Log("pass")
	} else {
		t.Error("hello fail")
	}
}
