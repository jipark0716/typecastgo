package main_test

import (
	"fmt"
	"testing"
	."github.com/jipark0716/typecastGo"
)

func TestMain(t *testing.T) {
	t.Error("1234")
	typecast, _ := NewTypeCast("jipark0716@gmail.com", "h4CZ2gDgHrydinP")
	for actor, _ := range Actor {
		blob, err := typecast.Exec("ㅋㅋ루삥뽕", actor)
		fmt.Printf("%s:%#v:%#v\n", actor, err, len(blob))
	}
}
