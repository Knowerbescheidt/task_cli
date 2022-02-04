/*
Copyright © 2022 Jan Olaf Eriksen HERE janolaferiksen@yahoo.de

*/
package main

import (
	"clitool/cmd"
	"clitool/db"
	"fmt"
	"os"
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	//Access denied error lösen über cmd mit admin rights
	// über main add task aufrufen und die builden mit go build
	//go install funktioniert aber das executable wird in gopath installiert also  C:\Users\janol\go\bin
	must(db.Init(dbPath))
	cmd.Execute()
	db.CloseDB()
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
