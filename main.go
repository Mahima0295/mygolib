package mygolib

import "log"

func LogInfo(message string) {
  log.Printf("INFO - %v", message)
}

func LogWarnig(message string) {
  log.Printf("Warn - %v", message)
}

func LogError(message string) {
  log.Printf("Error - %v", message)
}
