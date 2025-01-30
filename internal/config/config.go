package config

type config struct {
	dirPaths       []string
	dirPathFilters []string
}

var cfg = config{}

func SaveToDirPaths(dirPath string) {
	cfg.dirPaths = append(cfg.dirPaths, dirPath)
}

func LoadDirPaths() []string {
	return cfg.dirPaths
}

func SaveToDirPathFilters(dirPath string) {
	cfg.dirPathFilters = append(cfg.dirPathFilters, dirPath)
}

func LoadDirPathFilters() []string {
	return cfg.dirPathFilters
}
