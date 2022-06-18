package webserviceapp

import "net/http"
func RegisterRoutes(){
	http.HandleFunc("/login/google", HandleGoogleLogin)
	http.HandleFunc("/auth/google/callback", HandleGoogleCallback)
}


