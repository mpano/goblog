package cmd

import (
	"github.com/spf13/cobra"
	"goblog/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migration",
	Long:  `adding database migration`,
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	bootstrap.Migrate()
}
