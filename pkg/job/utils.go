package job

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// InitDotEnv loads .env file from the specified path.
func InitDotEnv(defaultPath string) error {
	path := defaultPath
	overridePath := os.Getenv("ENV_PATH")
	if overridePath != "" {
		path = os.Getenv("ENV_PATH")
	}

	err := godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	return nil
}

func GetArrayFrom(commaList string) []string {
	return strings.Split(commaList, ",")
}
