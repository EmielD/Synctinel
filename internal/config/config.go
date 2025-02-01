package config

type Config struct {
	dirPaths       []string
	dirPathFilters []string
}

var cfg = Config{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SaveToDirPaths(dirPath string) {
	cfg.dirPaths = append(cfg.dirPaths, dirPath)
}

func (c *Config) LoadDirPaths() []string {
	return cfg.dirPaths
}

func (c *Config) SaveToDirPathFilters(dirPath string) {
	cfg.dirPathFilters = append(cfg.dirPathFilters, dirPath)
}

func (c *Config) LoadDirPathFilters() []string {
	return cfg.dirPathFilters
}
