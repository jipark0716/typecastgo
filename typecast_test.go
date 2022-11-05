package typecastgo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	typecast := NewSession()
	err := typecast.Connect(&LoginRequest{
		Email:             os.Getenv("TYPECAST_EMAIL"),
		Password:          os.Getenv("TYPECAST_PASSWORD"),
		ReturnSecureToken: true,
	})
	if err != nil {
		log.Fatalf("123%+v", err)
	}

	request := NewRequest()
	request.Text = "ㅋㅋㄹ삥뽕"
	request.ActorId = "5c547544fcfee90007fed455"
	audio, err := typecast.Do([]*TypecastExecuteRequest{request})
	if err != nil {
		log.Fatalf("456%+v", err)
	}

	err = ioutil.WriteFile("./test.wav", audio, 0755)
	if err != nil {
		log.Fatalf("789%+v", err)
	}

}

func TestActor(t *testing.T) {
	typecast := NewSession()
	err := typecast.Connect(&LoginRequest{
		Email:             os.Getenv("TYPECAST_EMAIL"),
		Password:          os.Getenv("TYPECAST_PASSWORD"),
		ReturnSecureToken: true,
	})
	if err != nil {
		log.Fatalf("123%+v", err)
	}

	actors, err := typecast.GetActors()
	if err != nil {
		log.Fatalf("123%+v", err)
	}
	fmt.Printf("%#v\n", actors[0])
}
