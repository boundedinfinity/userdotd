package cmd

import (
	"fmt"

	"github.com/boundedinfinity/userdotd/model"
	"github.com/boundedinfinity/userdotd/system"
	"github.com/spf13/cobra"
)

func init() {
	for _, shell := range model.Shells {
		cmd := createShellCommand(shell)
		cmdStatus := createShellStatusCommand(shell)
		cmdInit := createShellInitCommand(shell)

		cmd.AddCommand(cmdInit)
		cmd.AddCommand(cmdStatus)
		shellCommand.AddCommand(cmd)
	}

	shellCommand.AddCommand(shellStatusCommand)
	RootCmd.AddCommand(shellCommand)
}

func createShellCommand(name string) *cobra.Command {
	return &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Manage %v shell", name),
	}
}

func createShellStatusCommand(name string) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: fmt.Sprintf("%v shell status", name),
		Run: func(cmd *cobra.Command, arg []string) {
			s := system.NewSystem(gContext, gLogger)
			status, err := s.ShellStatus(name)

			if err != nil {
				handleError(err)
			}

			fmt.Printf("%v: %v\n", status.Name, status.State)
		},
	}
}

func createShellInitCommand(name string) *cobra.Command {
	return &cobra.Command{
		Use:     "initialize",
		Aliases: []string{"init"},
		Short:   fmt.Sprintf("Initialize %v shell", name),
		Run: func(cmd *cobra.Command, arg []string) {
			s := system.NewSystem(gContext, gLogger)
			status, err := s.ShellInitialize(name)

			if err != nil {
				handleError(err)
			}

			fmt.Printf("%v: initialized\n", status.Name)
		},
	}
}

var shellCommand = &cobra.Command{
	Use:   "shell",
	Short: "Manage shells",
}

var shellStatusCommand = &cobra.Command{
	Use:   "status",
	Short: "Status for all shells",
	Run: func(cmd *cobra.Command, arg []string) {
		s := system.NewSystem(gContext, gLogger)
		statuses, err := s.ShellStatuses(model.Shells...)

		if err != nil {
			handleError(err)
		}

		for _, status := range statuses {
			fmt.Printf("%4v: %v\n", status.Name, status.State)
		}
	},
}
