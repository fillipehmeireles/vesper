package logUtils

import (
	"log"
)


func LogError(errorMsg string) {
	log.Printf("[ERROR] %s\n", errorMsg)
}

func LogInfo(infoMsg string) {
	log.Printf("[INFO] %s\n", infoMsg)
}
