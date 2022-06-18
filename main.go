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


type AppService struct{}

// Run method for provision service 
func (as *AppService) Run(){
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
  app:=AppService{}
  app.Run()
	fmt.Print("\n\n\n\t\t\tHi, I am Bihari123.Welcome to the game!\n\t\t\tThis is a project that I started for fun\n\n\nPlz enter your name to continue:\n")




}
