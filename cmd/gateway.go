/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"medica/microservices/api-gateway/processor"
	"medica/sdk/destination"

	"github.com/spf13/cobra"
)

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "api-gateway",
	Short: "Starter for api-gateway microservice",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		plantsAddress, _ := cmd.Flags().GetString("core-address")
		plantsPort, _ := cmd.Flags().GetString("core-port")
		plantsCfg := destination.NewConfig("core", plantsAddress, plantsPort)

		gatewayCfg := processor.Config{
			Plants: plantsCfg,
		}

		gatewayService := processor.NewDefaultGateway(gatewayCfg)

		gatewayService.Start()
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	gatewayCmd.PersistentFlags().String("core-port", "", "Port of plants service")
	gatewayCmd.PersistentFlags().String("core-address", "", "Address of plants service")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gatewayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
