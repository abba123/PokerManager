package oauth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const clientID string = "e0157b12f50fcc3b9b58"
const clientSecret string = "debdfae9b04bd9b2537249e53e3a1778650a19ad"
const scopes string = "user:email"

var Tokens map[string]string

func GenerateCodeURL() string {
	viper.AutomaticEnv()
	redirectURL := "http://" + viper.GetString("BACKEND") + ":8000/oauth/login"
	url := "https://github.com/login/oauth/authorize?client_id=%s&scope=%s&redirect_uri=%s"

	return fmt.Sprintf(url, clientID, scopes, redirectURL)
}

func GenerateTokenURL(code string) string {
	url := fmt.Sprintf("https://github.com/login/oauth/access_token?client_id=%s&client_secret=%s&code=%s", clientID, clientSecret, code)
	resp, err := http.Post(url, "application/json", nil)
	if err != nil{
		return ""
	}
	body, _ := ioutil.ReadAll(resp.Body)
	tokenString := strings.Split(string(body), "&")[0]
	token := strings.Split(tokenString, "=")[1]
	return token
}

func GetUser(token string) string {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://api.github.com/user", nil)

	req.Header = http.Header{
		"Authorization": []string{"token " + token},
	}

	res, _ := client.Do(req)

	var body struct {
		Username string `json:"login"`
	}

	r, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(r, &body)

	return body.Username
}

func StoreToken(IP string, token string) {
	Tokens[IP] = token
}

func CheckToken(IP string) string {
	tick := time.NewTicker(time.Second)
	token := ""
	exist := false
	for count := 0; count < 10; count++ {
		<-tick.C
		token, exist = Tokens[IP]
		if exist {
			break
		}
	}
	tick.Stop()
	return token
}
