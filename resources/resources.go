package resources

import (
	"fmt"
	"os"
)

type Resources struct {
	URL string
}

func NewResources() (*Resources, error) {
	sourceURL, err := findEnvVar("source_url")
	if err != nil {
		return nil, err
	}

	return &Resources{URL: sourceURL}, nil
}

func findEnvVar(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", fmt.Errorf("failed to find key %v in environment", key)
	}
	return value, nil
}
