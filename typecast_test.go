package typecast

import (
	"testing"
)

func TestMain(t *testing.T) {
	typecast := NewSession()
	typecast.Connect(&LoginRequest{
		Email:             "",
		Password:          "",
		ReturnSecureToken: true,
	})
}
