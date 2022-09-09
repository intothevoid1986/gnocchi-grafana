package utils

import "log"

// HandleError provide a standar error management
func HandleError(err error) {
	if err != nil {
		log.Panicf("Error: %v", err)
	}
}
