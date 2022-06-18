package config

type GoogleOauthConfig struct{
  ClientID string 
  ClientSecret string 
  Redirect string 
}

type OauthConfig struct{
  Google GoogleOauthConfig
}
