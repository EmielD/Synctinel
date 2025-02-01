package main

import (
	"github.com/emield/synctinel/internal/config"
	"github.com/emield/synctinel/internal/watcher"
)

func main() {
	config := config.NewConfig()

	config.SaveToDirPaths(".")
	config.SaveToDirPathFilters(".git")

	watcher.Init(config)
}
