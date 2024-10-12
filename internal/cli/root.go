package cli

import (
	"loader/internal/domain/usecase"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gitlab-downloader",
		Short: "Download gitlab repositories",
		Long:  `Download gitlab repositories`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	usecase.SaveConfig()
}

func AddCommand(cmd *cobra.Command) {
	rootCmd.AddCommand(cmd)
}

func AddCommands(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		AddCommand(cmd)
	}
}

func AddSubCommand(parent *cobra.Command, cmd *cobra.Command) {
	parent.AddCommand(cmd)
}

func AddSubCommands(parent *cobra.Command, cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		AddSubCommand(parent, cmd)
	}
}
