package main

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"

	"github.com/bihari123/terminal-based-game-in-go/config"
	"github.com/bihari123/terminal-based-game-in-go/dbmodel"
	"github.com/bihari123/terminal-based-game-in-go/helper"
	"github.com/bihari123/terminal-based-game-in-go/loghelper"
	"github.com/bihari123/terminal-based-game-in-go/services/web-service/webserviceapp"
)

var (
	DB *dbmodel.DB
)

var myConfig *config.Configuration

type AppService struct{}

// Run method for provision service
func (as *AppService) Run() {
	myConfig = helper.GetConfig()
	fmt.Printf("\n\n%+v\n\n", myConfig)
	fmt.Println("LogFolder --", myConfig.AppConfig.LogFolder)

	loghelper.ConfigLogurus(myConfig)
	setUpProcess()

	db, err := dbmodel.Init(myConfig)
	if err != nil {
		panic(err.Error())
	}
	DB = db

	sqlDb, _ := db.DB.DB()
	defer sqlDb.Close()

	webserviceapp.RegisterRoutes()
	webserviceapp.GetGooglOauthConfig(*myConfig)
 }

func setUpProcess() {
	runtime.GOMAXPROCS(2)
	debug.SetGCPercent(30)
}
func main() {
	app := AppService{}
	app.Run()
	fmt.Print("\n\n\n\t\t\tHi, I am Bihari123.Welcome to the game!\n\t\t\tThis is a project that I started for fun\n\n\nPlz enter your name to continue:  ")
	//
	// var name string
	// fmt.Scanln(&name)
	//
	// var email string
	// fmt.Print("Plz enter your google email:  ")
	// fmt.Scanln(&email)
	//
	// // remember asking password from user in your software is not a good practice, but we do it anyway
	// var password string
	// fmt.Print("Plz enter your password")
	// fmt.Scanln(&password)

	helper.OpenUrl("http://localhost:8086/login/google")
 
  http.ListenAndServe(":8086",nil)

}
