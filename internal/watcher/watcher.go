package watcher

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/emield/synctinel/internal/config"
	"github.com/fsnotify/fsnotify"
)

var dirPathFilters = []string{}

func Init(config *config.Config) {

	dirPathFilters = config.LoadDirPathFilters()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	for _, path := range config.LoadDirPaths() {
		subDirs, err := WalkDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, subDir := range subDirs {
			if subDir == "." {
				continue
			}

			err = watcher.Add(subDir)
			if err != nil {
				log.Fatal(err)
			}
		}

		err = watcher.AddWith(path)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Start listening for events
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	<-make(chan struct{})
}

func WalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && pathNotInFilters(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, err
}

func pathNotInFilters(path string) bool {
	for _, filter := range dirPathFilters {
		if strings.Contains(path, filter) {
			return false
		}
	}
	return true
}
