package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"loader/internal/domain/usecase"
	"loader/internal/infrastructure/formating"
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

func registerSetTokenCmd() *cobra.Command {
	var (
		set string
	)

	tokenCmd := &cobra.Command{
		Use:   "set-token",
		Short: "Установить токен",
		Long:  "Установить токен для доступа к GitLab API и репозиториям. ",
		Run: func(cmd *cobra.Command, args []string) {
			usecase.SetToken(set)
		},
	}

	tokenCmd.Flags().StringVarP(&set, "token", "t", "", "Передать токен для установки")

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
				"По умолчанию используется %s\n"+
				"При необходимости можно изменить на другой адрес командой %s\n",
			formating.YellowBoldText(viper.Get("addr").(string)), formating.YellowBoldText("gitlab-downloader addr -s <адрес>")),

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
		registerSetTokenCmd(),
		registerRepoCmd(),
	)
}
