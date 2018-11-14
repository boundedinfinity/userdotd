package cmd

import (
	"log"

	"github.com/boundedinfinity/userdotd/system"
	"github.com/spf13/cobra"
)

var DisabledCmd = &cobra.Command{
	Use:   "disabled",
	Short: "disabled",
	Long:  "disabled",
}

var DisabledLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long:  "login",
	Run: func(cmd *cobra.Command, args []string) {
		s := system.NewSystem(gContext, gLogger)

		if len(args) > 0 {
			if err := s.Disable(args[0]); err != nil {
				log.Fatal(err.Error())
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(DisabledCmd)
	DisabledCmd.AddCommand(DisabledLoginCmd)
}
