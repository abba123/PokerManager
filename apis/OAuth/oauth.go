package oauth

import "fmt"

const clientID string = "2cab12bd0bcae1f150d0"
const clientSecret string = "8517b74a34215561a21abae585cfb2ac4f0fa1be"
const scopes string = "user:email"
const redirectURL string = "http://127.0.0.1/oauth/login"

func GenerateURL() string {
	url := "https://github.com/login/oauth/authorize?client_id=%s&scope=%s&redirect_uri=%s"

	return fmt.Sprintf(url, clientID, scopes, redirectURL)
}
