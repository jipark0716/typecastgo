package typecast

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestMain(t *testing.T) {
	typecast := NewSession()
	err := typecast.Connect(&LoginRequest{
		Email:             "email",
		Password:          "password",
		ReturnSecureToken: true,
	})
	if err != nil {
		fmt.Printf("123%+v", err)
	}

	request := NewRequest()
	request.Text = "ㅋㅋㄹ삥뽕"
	request.ActorId = "5c547544fcfee90007fed455"
	audio, err := typecast.Do([]*TypecastExecuteRequest{request})
	if err != nil {
		fmt.Printf("456%+v", err)
	}

	println(typecast.Token.IdToken)
	err = ioutil.WriteFile("./test.wav", audio, 0755)
	if err != nil {
		fmt.Printf("789%+v", err)
	}

}
