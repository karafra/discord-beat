// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period         time.Duration `config:"period"`
	ApiToken       string        `config:"apiToken"`
	ExportGuilds   []string      `config:"exportGuilds"`
	ExportChannels []string      `config:"exportChannels"`
}

var DefaultConfig = Config{
	Period:         1 * time.Second,
	ApiToken:       "",
	ExportGuilds:   make([]string, 0),
	ExportChannels: make([]string, 0),
}
