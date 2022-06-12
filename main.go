package main

import (
	"fmt"
	"runtime"
	"runtime/debug"

	"github.com/bihari123/terminal-based-game-in-go/dbmodel"
	"github.com/bihari123/terminal-based-game-in-go/helper"
	"github.com/bihari123/terminal-based-game-in-go/loghelper"
)

var(
	DB *dbmodel.DB  
)

// ProvisionService - Struct for provisioning service
type ProvisionService struct{}

// Run method for provision service 
func (ps *ProvisionService) Run(){
  myConfig:= helper.GetConfig()
  fmt.Printf("\n\n%+v\n\n",myConfig)
  fmt.Println("LogFolder --",myConfig.AppConfig.LogFolder)

  loghelper.ConfigLogurus(myConfig)
  setUpProcess()

  db,err:=dbmodel.Init(myConfig)
  if err!=nil{
  	panic(err.Error())
  }
  DB=db

  sqlDb,_:=db.DB.DB()
  defer sqlDb.Close()
}

func setUpProcess(){
  runtime.GOMAXPROCS(2)
  debug.SetGCPercent(30)
}
func main() {
  app:=ProvisionService{}
  app.Run()
}
