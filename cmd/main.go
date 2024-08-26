package main

import (
	"fmt"
	"go-auth/internal/config"
	database "go-auth/internal/db"
	"go-auth/internal/server"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

func main() {
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	go func() {
		log.Println("Start pprof...")
		log.Println(http.ListenAndServe(":6060", nil))
	}()

	c := config.NewConfig()

	db := database.NewDatabase(c)
	db.Connect()

	server.NewServer(c, db).Start()
}
