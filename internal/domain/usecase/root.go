package usecase

import "loader/internal/infrastructure/viper"

func SaveConfig() {
	viper.SaveConfig()
}
