package cmd

import (
	"log"

	"github.com/boundedinfinity/userdotd/system"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialization",
	Long:  "initialization",
	Run: func(cmd *cobra.Command, arg []string) {
		s := system.NewSystem(gContext, gLogger)

		if err := s.InitializeSystem(); err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(InitCmd)
}
