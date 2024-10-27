package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"loader/internal/domain/entity"
	"loader/internal/domain/usecase"
	"loader/internal/infrastructure/formating"
	"loader/internal/infrastructure/viper"
	"strings"
)

func registerConfigCmd() *cobra.Command {
	helpMsg := fmt.Sprintf(`
Управление настройками для доступа к GitLab API и репозиториям.

Текущие настройки:
%s`,
		formating.MapToString(entity.CFG))

	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Управление настройками",
		Long:  helpMsg,
	}

	return configCmd
}

func registerTokenCmd() *cobra.Command {
	helpMsg := fmt.Sprintf(`
Сохраненный токен: %s

Для получения токена перейди по ссылке: %s
При создании токена необходимо дать ему права на чтение репозиториев
После создания токена его можно сохранить с помощью команды %s или использовать напрямую
командой %s
`,
		formating.YellowBoldText(viper.Get("token").(string)),
		formating.YellowBoldText(fmt.Sprintf("%s/profile/personal_access_tokens", strings.Replace(entity.CFG["addr"].(string), "api/v4", "-", 1))),
		formating.YellowBoldText("gitlab-downloader config token -s <токен>"),
		formating.YellowBoldText("gitlab-downloader download -t <токен>"),
	)
	var (
		set string
	)

	tokenCmd := &cobra.Command{
		Use:   "token",
		Short: "Управление токеном",
		Long:  helpMsg,
		Run: func(cmd *cobra.Command, args []string) {
			if set != "" {
				usecase.SetToken(set)
			}
		},
	}
	tokenCmd.Flags().StringVarP(&set, "set", "s", "", "Сохранить токен в конфиг")

	return tokenCmd
}

func registerRepoCmd() *cobra.Command {
	var (
		set string
	)

	addrCmd := &cobra.Command{
		Use:   "addr",
		Short: "Управление сервером",
		Long: fmt.Sprintf(
			"Управление сервером для доступа к GitLab API и репозиториям.\n"+
				"Сервер представляет собой URL GitLab сервера, к которому будет производиться доступ. "+
				"По умолчанию используется %s\n"+
				"При необходимости можно изменить на другой адрес командой %s\n",
			formating.YellowBoldText(viper.Get("addr").(string)), formating.YellowBoldText("gitlab-downloader addr -s <адрес>")),

		Run: func(cmd *cobra.Command, args []string) {
			if set != "" {
				usecase.SetAddr(set)
			}
		},
	}

	addrCmd.Flags().StringVarP(&set, "set", "s", "", "Задать целевой сервер")

	return addrCmd
}

func init() {
	configCmd := registerConfigCmd()

	AddSubCommands(
		configCmd,
		registerTokenCmd(),
		registerRepoCmd(),
	)

	AddCommands(
		configCmd,
	)
}
