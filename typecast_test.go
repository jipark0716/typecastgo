package typecast_test

import (
	"fmt"
	"testing"
	"os"
	."github.com/jipark0716/typecastGo"
)

func TestMain(t *testing.T) {
	t.Error("1234")
	typecast, _ := NewTypeCast(os.Getenv("TYPECAST_ID"), os.Getenv("TYPECAST_pw"))
	for actor, _ := range Actor {
		go func(actor string) {
			blob, err := typecast.Exec("ㅋㅋ루삥뽕", actor)
			fmt.Printf("%s:%#v:%#v\n", actor, err, len(blob))
		}(actor)
	}
	select{}
}
