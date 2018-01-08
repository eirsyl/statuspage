package cmd

import (
	"github.com/spf13/cobra"
)

func init() {

}

var RootCmd = &cobra.Command{
	Use:   "statuspage",
	Short: "Statuspage platform written in golang with postgres as the backing datastore.",
	Long:  ``,
}
