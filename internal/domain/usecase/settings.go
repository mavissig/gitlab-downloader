package usecase

import (
	"fmt"
	"loader/internal/infrastructure/formating"
	"loader/internal/infrastructure/viper"
)

func SetToken(val string) {
	viper.Set("token", val)
	fmt.Println(formating.LogSuccess("Токен установлен"))
}

func ShowRepoInfo() {

}
