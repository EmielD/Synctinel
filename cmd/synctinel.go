package main

import (
	"fmt"

	"github.com/emield/synctinel/internal/config"
	"github.com/emield/synctinel/internal/watcher"
)

func main() {
	config.SaveToDirPaths(".")
	config.SaveToDirPathFilters("internal")
	config.SaveToDirPathFilters(".git")

	fmt.Println(len(config.LoadDirPathFilters()))

	watcher.Init()
}
