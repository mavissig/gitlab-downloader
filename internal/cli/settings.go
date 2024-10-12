package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"loader/internal/domain/usecase"
	"loader/internal/infrastructure/formatting"
	"loader/internal/infrastructure/viper"
)

func registerTokenCmd() *cobra.Command {
	var (
		info bool
		set  string
	)

	tokenCmd := &cobra.Command{
		Use:   "token",
		Short: "Управление токеном",
		Long:  "Управление токеном для доступа к GitLab API и репозиториям. ",
		Run: func(cmd *cobra.Command, args []string) {
			if info {
				usecase.ShowTokenInfo()
			}
			if set != "" {
				usecase.SetToken(set)
			}
		},
	}

	tokenCmd.Flags().BoolVarP(&info, "info", "i", false, "Показать токен и дополнительную информацию")
	tokenCmd.Flags().StringVarP(&set, "set", "s", "", "Установить токен")

	return tokenCmd
}

func registerRepoCmd() *cobra.Command {
	var (
		show bool
		set  string
	)

	tokenCmd := &cobra.Command{
		Use:   "addr",
		Short: "Управление сервером",
		Long: fmt.Sprintf(
			"Управление сервером для доступа к GitLab API и репозиториям.\n"+
				"Сервер представляет собой URL GitLab сервера, к которому будет производиться доступ. "+
				"По умолчанию используется %s%s%s\n"+
				"При необходимости можно изменить на другой адрес командой %sgitlab-downloader addr -s <адрес>%s\n",
			formatting.YELLOW_BOLD, viper.Get("addr"), formatting.NC, formatting.YELLOW_BOLD, formatting.NC),

		Run: func(cmd *cobra.Command, args []string) {
			if show {
				usecase.ShowTokenInfo()
			}
			if set != "" {
				usecase.SetToken(set)
			}
		},
	}

	tokenCmd.Flags().BoolVarP(&show, "info", "i", false, "Показать целевой сервер и дополнительную информацию")
	tokenCmd.Flags().StringVarP(&set, "set", "s", "", "Задать целевой сервер")

	return tokenCmd
}

func init() {
	AddCommands(
		registerTokenCmd(),
		registerRepoCmd(),
	)
}
