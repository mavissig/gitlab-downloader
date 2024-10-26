package cli

import (
	"fmt"
	"loader/internal/domain/usecase"
	"loader/internal/infrastructure/formating"
	"log"
	"os"

	"github.com/spf13/cobra"
)

/*
   ---------------------------------------------
				     MESSAGES
   ---------------------------------------------
*/

var (
	rootCmdLongMsg = fmt.Sprintf(`
Утилита позволяет скачивать репозитории с GitLab по заданным настройкам

Пример использования: %s

Для получения подробной информации по настройкам воспользуйтесь командой: %s
`,
		formating.YellowBoldText("gitlab-downloader download -t <token>"),
		formating.YellowBoldText("gitlab-downloader config --help"),
	)
)

/*
	---------------------------------------------
					  COMMANDS
	---------------------------------------------
*/

var (
	rootCmd = &cobra.Command{
		Use:   "gitlab-downloader",
		Short: "Утилита для скачивания репозиториев с GitLab",
		Long:  rootCmdLongMsg,
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
