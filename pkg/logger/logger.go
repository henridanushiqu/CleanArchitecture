package logger

import (
	"context"
	"log"
)

func LogDebug(ctx context.Context, message string) {
	log.Printf("[DEBUG] %s", message)
}

func LogError(ctx context.Context, message string, err error) {
	log.Printf("[ERROR] %s: %v", message, err)
}
