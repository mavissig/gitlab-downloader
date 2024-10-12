package utils

import (
	"log"
	"os"
	"strings"
)

func CreateChannel[t any]() chan t {
	return make(chan t, 10)
}

func CreateChannelPair[t any]() (in chan<- t, out <-chan t) {
	ch := make(chan t, 10)
	return ch, ch
}

func NormalizePath(path string) string {
	if path == "" {
		return ""
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("[UTILS][NormalizePath] ERROR: ", err)
		return path
	}

	path = strings.Replace(path, "~", userHomeDir, -1)

	return path
}
