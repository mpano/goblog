package bootstrap

import (
	"goblog/internal/database/migration"
	"goblog/pkg/config"
	"goblog/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()

}
