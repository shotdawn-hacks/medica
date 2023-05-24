package config

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func ReadAllAndUnmarshalYAML(source io.Reader, v interface{}) error {
	b, err := io.ReadAll(source)
	if err != nil {
		return fmt.Errorf("source unmarshaling failed: %w", err)
	}

	err = yaml.Unmarshal(b, v)
	if err != nil {
		return fmt.Errorf("source unmarshaling failed: %w", err)
	}

	return nil
}

func ReadFileAndUnmarshalYAML(path string, v interface{}) error {
	reader, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open failed: %w", err)
	}

	return ReadAllAndUnmarshalYAML(reader, v)
}
