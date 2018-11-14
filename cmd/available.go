package cmd

import (
	"log"

	"github.com/boundedinfinity/userdotd/system"
	"github.com/spf13/cobra"
)

var AvailableCmd = &cobra.Command{
	Use:   "available",
	Short: "available",
	Long:  "available",
}

var AvailableLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long:  "login",
	Run: func(cmd *cobra.Command, arg []string) {
		s := system.NewSystem(gContext, gLogger)

		if err := s.AvailableLogin(); err != nil {
			log.Fatal(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(AvailableCmd)
	AvailableCmd.AddCommand(AvailableLoginCmd)
}
