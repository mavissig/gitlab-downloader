package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"loader/internal/domain/entity"
	"loader/internal/domain/usecase"
)

func registerDownloadCmd() *cobra.Command {
	msg := `
Скачать проекты по заданным настройкам

Предустановленные настройки можно посмотреть командой:
gitlab-downloader config [ --help | -h ]

Если ты студент школы 21, то тебе необходимо только установить токен, а остальные настройки уже установлены.

Так же можно воспользоваться утилитой без сохранения настроек, передав токен напрямую:
gitlab-downloader download -t <token>
`
	downloadCmd := &cobra.Command{
		Use:   "download",
		Short: "Скачать проекты",
		Long:  fmt.Sprintf(msg),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(entity.CFG)
			return usecase.DownloadValidate()
		},
		Run: func(cmd *cobra.Command, args []string) {
			usecase.DownloadProjects()
		},
	}

	downloadCmd.Flags().StringVarP(&entity.TOKEN, "token", "t", "", "Токен для доступа к GitLab API")

	return downloadCmd
}

func init() {
	AddCommands(
		registerDownloadCmd(),
	)
}
