package cmd

import (
	"log"

	"github.com/boundedinfinity/userdotd/system"
	"github.com/spf13/cobra"
)

var EnabledCmd = &cobra.Command{
	Use:   "enabled",
	Short: "enabled",
	Long:  "enabled",
}

var EnabledLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long:  "login",
	Run: func(cmd *cobra.Command, args []string) {
		s := system.NewSystem(gContext, gLogger)

		if len(args) > 0 {
			if err := s.Enable(args[0]); err != nil {
				log.Fatal(err.Error())
			}
		} else {
			if err := s.Enabled(); err != nil {
				log.Fatal(err.Error())
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(EnabledCmd)
	EnabledCmd.AddCommand(EnabledLoginCmd)
}
