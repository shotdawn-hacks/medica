/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"medica/microservices/core/processor"

	"github.com/spf13/cobra"
)

// gatewayCmd represents the gateway command
var coreCmd = &cobra.Command{
	Use:   "core",
	Short: "Starter for core microservice",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		coreService := processor.NewDefaultCore()

		coreService.Start()
	},
}

func init() {
	rootCmd.AddCommand(coreCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gatewayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
