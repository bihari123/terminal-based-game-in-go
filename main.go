package main

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/bihari123/terminal-based-game-in-go/helper"
)

// ProvisionService - Struct for provisioning service
type ProvisionService struct{}

// Run method for provision service 
func (ps *ProvisionService) Run(){
  myConfig:= helper.GetConfig()
  fmt.Println("LogFolder --",myConfig.AppConfig.LogFolder)
  setUpProcess()
}

func setUpProcess(){
  runtime.GOMAXPROCS(2)
  debug.SetGCPercent(30)
}
func main() {
  app:=ProvisionService{}
  app.Run()
}
