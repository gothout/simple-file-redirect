package env

import (
	"fmt"
	envApp "simple-file-redirect/internal/configuration/env/application"
	envServer "simple-file-redirect/internal/configuration/env/server"

	"github.com/joho/godotenv"
)

// CheckEnvs checks if the .env file exists and returns an appropriate error if it doesn't.
func CheckEnvs() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// Validate server envs
	err = envServer.ValidateServerEnv()
	if err != nil {
		return err
	}
	// Validate application envs
	err = envApp.ValidateApplicationEnv()
	if err != nil {
		return err
	}

	return nil
}
