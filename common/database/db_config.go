package database

type Database struct {
	MaxIdle      int               `yaml:"max_idle"`
	MaxConn      int               `yaml:"max_conn"`
	MaxLifeCycle int               `yaml:"max_life_cycle"`
	LogLevel     string            `yaml:"log_level"`
	Databases    map[string]string `yaml:"databases"`
}
