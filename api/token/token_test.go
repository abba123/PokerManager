package token

import (
	"testing"
)

func TestToken(t *testing.T) {
	token := GenerateToken("test")
	result, err := ValidToken(token)

	if err == nil && result.Username == "test"{
		t.Log("Token PASS")
	}else{
		t.Log("Token FAIL")
	}
}
