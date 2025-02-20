package main

import (
	"github.com/henrriusdev/portfolio/api"
	"github.com/henrriusdev/portfolio/config"
	"github.com/henrriusdev/portfolio/pkg/store"
)

func main() {
	config.LoadEnv()
	db, _ := store.InitDatabase()
	api.Start(db)
}
