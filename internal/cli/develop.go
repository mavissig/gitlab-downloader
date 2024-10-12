package cli

import (
	"github.com/spf13/cobra"
	"loader/internal/domain/usecase"
)

var (
	developCmd = &cobra.Command{
		Use: "dev",
		Run: func(cmd *cobra.Command, args []string) {
			info, _ := cmd.Flags().GetBool("info")
			if info {
				usecase.ShowInfo()
			}
			config, _ := cmd.Flags().GetBool("config")
			if config {
				usecase.ShowConfig()
			}
		},
	}
)

func init() {
	AddCommands(
		developCmd,
	)

	developCmd.Flags().BoolP("info", "i", false, "Show info")
	developCmd.Flags().BoolP("config", "c", false, "Show config")
}
