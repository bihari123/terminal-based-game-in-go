package main

import (
	"net/http"

	"github.com/bihari123/terminal-based-game-in-go/loghelper"
	"github.com/gorilla/mux"
)

func main(){
  r:=mux.NewRouter()
  RegisterRoutes(r)
  loghelper.Log("starting the server at 8082")
  http.ListenAndServe(":8082",nil)
}
