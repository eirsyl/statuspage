package cli

import (
	"github.com/spf13/cobra"
)

func init() {
}

var IncidentCmd = &cobra.Command{
	Use:   "incident",
	Short: "Manage incidents and updates",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
