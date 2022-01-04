package token

import (
	"fmt"
	"testing"
)

func TestToken(t *testing.T) {
	token := GenerateToken("test")
	result, err := ValidToken(token)
	fmt.Println(result)
	if err == nil && result.Username == "test" {
		t.Log("Token PASS")
	} else {
		t.Error("Token FAIL")
	}
}

func TestTokenErr(t *testing.T) {
	token := GenerateToken("test")
	token += " "
	_, err := ValidToken(token)

	if err != nil {
		t.Log("TokenErr PASS")
	} else {
		t.Error("Token Failed FAIL")
	}
}
