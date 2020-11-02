package cmd

import (
	"github.com/idasilva/project-web/app"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Create a new sample server",
	Long:  "Long description",
	RunE: func(cmd *cobra.Command, args []string) error {

		app := app.Run()
		if err := app.Serve(); err != nil {
			return err
		}

		return nil

	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

}
