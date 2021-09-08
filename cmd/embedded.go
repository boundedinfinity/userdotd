package cmd

import (
	"fmt"

	"github.com/boundedinfinity/userdotd/model"
	"github.com/boundedinfinity/userdotd/system"
	"github.com/spf13/cobra"
)

func init() {
	embeddedListCommand.AddCommand(embeddedListFilesCommand)
	embeddedListCommand.AddCommand(embeddedListDirsCommand)
	embeddedListCommand.AddCommand(embeddedListAllCommand)
	embeddedCommand.AddCommand(embeddedListCommand)
	RootCmd.AddCommand(embeddedCommand)
}

var embeddedCommand = &cobra.Command{
	Use:   "embedded",
	Short: "Query embedded resources",
}

var embeddedListCommand = &cobra.Command{
	Use:   "list",
	Short: "list embedded resources",
}

type listFn func() (model.EmbeddedListResponse, error)

func createListCommand(use, short string, aliases []string, fn listFn) *cobra.Command {
	return &cobra.Command{
		Use:     use,
		Short:   short,
		Aliases: aliases,
		Run: func(cmd *cobra.Command, arg []string) {
			response, err := fn()

			if err != nil {
				handleError(err)
			}

			for _, file := range response.Paths {
				fmt.Printf("%v\n", file)
			}
		},
	}
}

var embeddedListFilesCommand = createListCommand(
	"files", "list embedded files", nil,
	func() (model.EmbeddedListResponse, error) {
		s := system.NewSystem(gContext, gLogger)
		return s.EmbeddedListFiles(model.EmbeddedListRequest{})
	},
)

var embeddedListAllCommand = createListCommand(
	"all", "list embedded files and directories", nil,
	func() (model.EmbeddedListResponse, error) {
		s := system.NewSystem(gContext, gLogger)
		return s.EmbeddedListAll(model.EmbeddedListRequest{})
	},
)

var embeddedListDirsCommand = createListCommand(
	"directories", "list embedded directories", []string{"dirs"},
	func() (model.EmbeddedListResponse, error) {
		s := system.NewSystem(gContext, gLogger)
		return s.EmbeddedListDirs(model.EmbeddedListRequest{})
	},
)
