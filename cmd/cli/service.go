package cli

import (
	"github.com/eirsyl/statuspage/pkg/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
)

func init() {
	ServiceCmd.AddCommand(listServiceCmd)
	ServiceCmd.AddCommand(createServiceCmd)
	ServiceCmd.AddCommand(updateServiceCmd)
	ServiceCmd.AddCommand(deleteServiceCmd)
	ServiceCmd.AddCommand(getServiceCmd)
}

var ServiceCmd = &cobra.Command{
	Use:   "service",
	Short: "Manage services",
}

var listServiceCmd = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		apiUrl, token := viper.GetString("apiUrl"), viper.GetString("token")
		api, err := api.NewAPI(apiUrl, token)
		if err != nil {
			log.Fatal(err)
		}
		services, err := api.ListServices()
		log.Info(services, err)
	},
}

var createServiceCmd = &cobra.Command{
	Use: "create",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var updateServiceCmd = &cobra.Command{
	Use:  "update [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var deleteServiceCmd = &cobra.Command{
	Use:  "delete [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var getServiceCmd = &cobra.Command{
	Use:  "get [id]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	},
}
