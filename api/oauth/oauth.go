package oauth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const clientID string = "2cab12bd0bcae1f150d0"
const clientSecret string = "8517b74a34215561a21abae585cfb2ac4f0fa1be"
const scopes string = "user:email"

const redirectURL string = "http://localhost:8000/oauth/login"
//const redirectURL string = "http://3.133.150.55/oauth/login"

var OAuthChan chan string

func GenerateCodeURL() string {
	url := "https://github.com/login/oauth/authorize?client_id=%s&scope=%s&redirect_uri=%s"

	return fmt.Sprintf(url, clientID, scopes, redirectURL)
}

func GenerateTokenURL(code string) string {
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	resp, _ := http.Post(url, "application/json", nil)
	body, _ := ioutil.ReadAll(resp.Body)

	tokenString := strings.Split(string(body), "&")[0]
	token := strings.Split(tokenString, "=")[1]

	fmt.Println(token)
	return token
}