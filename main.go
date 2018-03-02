package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"mytank/rest"
)

func main() {
	rest.PrepareConfigs()
	context := rest.NewContext()
	defer context.Destroy()

	http.Handle("/",context.Router)
	dotPort := fmt.Sprintf(":%v",rest.CONFIG.ServerPort)

	info := fmt.Sprintf("App started at http://localhost%v", dotPort)
	rest.LogInfo(info)

	err := http.ListenAndServe(dotPort,nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}