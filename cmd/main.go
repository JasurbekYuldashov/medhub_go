package main

import (
	"github.com/jasurbekyuldashov/medhub_go/db"
	"github.com/jasurbekyuldashov/medhub_go/routes"
	"log"

	_ "github.com/lib/pq"
)

func main() {

	log.Println("Postgres Connection successfully done!")

	db.Init()
	routes.Run()
}
