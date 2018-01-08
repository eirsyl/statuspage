package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/eirsyl/statuspage/cmd/cli"
)

func init() {
	cliCmd.PersistentFlags().StringP("apiUrl", "a", "http://127.0.0.1:8080", "The URL used to access the API")
	cliCmd.PersistentFlags().StringP("token", "t", "", "The token used for authorizing with the API")
	viper.BindPFlag("apiUrl", cliCmd.PersistentFlags().Lookup("apiUrl"))
	viper.BindPFlag("token", cliCmd.PersistentFlags().Lookup("token"))
	viper.BindEnv("apiUrl", "API_URL")
	viper.BindEnv("token", "TOKEN")
	cliCmd.AddCommand(cli.IncidentCmd)
	cliCmd.AddCommand(cli.ServiceCmd)
	RootCmd.AddCommand(cliCmd)
}

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Access the statuspage api from the command line",
}
