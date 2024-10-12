package usecase

import (
	"fmt"
	"loader/internal/infrastructure/formating"
	"loader/internal/infrastructure/viper"
)

func ShowTokenInfo() {
	fmt.Println("Твой токен: ", viper.Get("token"))
	fmt.Println()
	fmt.Println("Для получения токена перейди по ссылке: \033[0;33m https://gitlab.com/profile/personal_access_tokens\033[0m")
	fmt.Println("При создании токена необходимо дать ему права на чтение репозиториев")
	fmt.Println("После создания токена необходимо установить его с помощью команды \033[0;33m gitlab-downloader token -s <токен>\033[0m")
}

func SetToken(val string) {
	viper.Set("token", val)
	fmt.Println(formating.LogSuccess("Токен установлен"))
}

func ShowRepoInfo() {

}
