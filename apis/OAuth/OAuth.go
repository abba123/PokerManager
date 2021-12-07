package OAuth

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const clientID string = "519659208785-fdr3eog5vs6krnqfbeigev8matphkl85.apps.googleusercontent.com"
const clientSecert string = "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"

//const redirectUrl string = "127.0.0.1/oauth/login/"

const redirectUrl string = "3.133.150.55/oauth/login/"

// oauthClient shows how to use an OAuth client ID to authenticate as an end-user.
func GetOAuthUrl() string {

	response_type := "code"
	scope := "https://www.googleapis.com/auth/userinfo.profile"

	url := "https://accounts.google.com/o/oauth2/v2/auth?client_id=%s&response_type=%s&scope=%s&redirect_uri=%s"

	return fmt.Sprintf(url, clientID, response_type, scope, redirectUrl)
}

func GetOAuthToken(code string) {
	//token, err := accessToken(code)

}

func accessToken(code string) (string, error) {
	u := "https://www.googleapis.com/oauth2/v4/token"

	data := url.Values{"code": {code}, "client_id": {clientID}, "client_secret": {clientSecert}, "grant_type": {"authorization_code"}, "redirect_uri": {redirectUrl}}
	body := strings.NewReader(data.Encode())

	var token string

	resp, err := http.Post(u, "application/x-www-form-urlencoded", body)
	if err != nil {
		return token, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	fmt.Println(b)
	//token = gjson.GetBytes(b, "access_token").String()

	return token, nil
}
