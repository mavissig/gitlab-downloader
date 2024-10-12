package usecase

import (
	"fmt"
	"loader/internal/domain/entity"
	"log"
	"os"
)

func ShowInfo() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Println("[UC][Homedir] ERROR: ", err)
		return
	}

	executable, err := os.Executable()
	if err != nil {
		log.Println("[UC][Executable] ERROR: ", err)
		return
	}

	fmt.Println("Home directory: ", homedir)
	fmt.Println("Current directory: ", os.Args[0])
	fmt.Println("Executable: ", executable)

	fmt.Println("config: ", fmt.Sprintf("%s/.config/", homedir))
}

func ShowConfig() {
	fmt.Println("config: ", entity.CFG)
}
