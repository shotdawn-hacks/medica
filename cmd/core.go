/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"medica/microservices/core/processor"
	"medica/sdk/db"

	"github.com/spf13/cobra"
)

// gatewayCmd represents the gateway command
var coreCmd = &cobra.Command{
	Use:   "core",
	Short: "Starter for core microservice",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		postgresURI, _ := cmd.Flags().GetString("postgres")
		db.SetURI(postgresURI)

		coreService := processor.NewDefaultCore()

		coreService.Start()
	},
}

func init() {
	rootCmd.AddCommand(coreCmd)

	coreCmd.PersistentFlags().String("postgres", "", "PostgreSQL connection string")
}
