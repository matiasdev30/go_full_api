package main

import (
	"github.com/matiasdev30/go_api/config"
	"github.com/matiasdev30/go_api/db"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main(){
	db.InitDatabase()
	config.InitGin()
}