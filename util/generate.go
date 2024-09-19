package util

import (
	"crypto/rand"
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

// GenerateUniqueFileName generates a unique file name using a timestamp and a random number.
func GenerateUniqueFileName(originalName string) string {
	// Get the current timestamp
	now := time.Now().Format("20060102150405")

	// Generate a random 6-byte number
	randomNum := make([]byte, 6)
	_, err := rand.Read(randomNum)
	if err != nil {
		panic(err) // Handle error appropriately in production code
	}
	uniqueID := fmt.Sprintf("%x", randomNum)

	// Get file extension
	ext := filepath.Ext(originalName)
	baseName := strings.TrimSuffix(originalName, ext)

	// Create the unique file name
	return fmt.Sprintf("%s-%s-%s%s", now, uniqueID, baseName, ext)
}
