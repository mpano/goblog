package config

import "goblog/config"

func Get() config.Config {
	return configuration
}
