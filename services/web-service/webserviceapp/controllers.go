package webserviceapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bihari123/terminal-based-game-in-go/apiresponse"
	"github.com/bihari123/terminal-based-game-in-go/config"
	"github.com/bihari123/terminal-based-game-in-go/loghelper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)


var (
	googleOauthConfig *oauth2.Config
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

func GetGooglOauthConfig(myConfig config.Configuration){
  googleOauthConfig = &oauth2.Config{
		RedirectURL:  myConfig.Oauth.Google.Redirect,
		ClientID:     myConfig.Oauth.Google.ClientID,
		ClientSecret: myConfig.Oauth.Google.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	content, err := getUserInfo(r.FormValue("state"), r.FormValue("code"))
	if err != nil {
		fmt.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
  resp:=apiresponse.Success(http.StatusFound,"user verified",content)
  loghelper.Log(fmt.Sprintf("\n\n\n%+v\n\n\n",resp))
  json.NewEncoder(w).Encode(resp)
  return
	
}

func getUserInfo(state string, code string) ([]byte, error) {
	if state != oauthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}

	return contents, nil
}
