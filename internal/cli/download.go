package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"loader/internal/domain/entity"
	"loader/internal/domain/usecase"
)

func registerDownloadCmd() *cobra.Command {
	tokenCmd := &cobra.Command{
		Use:   "download",
		Short: "Скачать проекты",
		Long:  "Скачать проекты по заданным настройкам",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(entity.CFG)
			return usecase.DownloadValidate()
		},
		Run: func(cmd *cobra.Command, args []string) {
			usecase.DownloadProjects()
		},
	}

	return tokenCmd
}

func init() {
	AddCommands(
		registerDownloadCmd(),
	)
}
