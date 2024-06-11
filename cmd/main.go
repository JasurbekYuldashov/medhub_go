package main

import (
	"github.com/jasurbekyuldashov/medhub_go/db"
	"github.com/jasurbekyuldashov/medhub_go/routes"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	db.Init()
	log.Println("Postgres Connection successfully done!")

	db.InitRedis(1)
	log.Println("Redis Connection successfully done!")

	routes.Run()
}
