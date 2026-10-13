package generator

import (
	"embed"
	"fmt"

	"go.yaml.in/yaml/v3"
)

// TenseData contains regular endings and per-verb conjugation overrides.
type TenseData struct {
	Name           string                       `yaml:"name"`
	RegularEndings map[string]map[string]string `yaml:"regular_endings"`
	Irregular      map[string]map[string]string `yaml:"irregular"`
}

//go:embed data/*.yml
var verbFiles embed.FS

// LoadTenseData reads an embedded YAML file containing regular endings and irregular conjugations.
func LoadTenseData(filename string) (*TenseData, error) {
	data, err := verbFiles.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", filename, err)
	}

	var tense TenseData
	if err := yaml.Unmarshal(data, &tense); err != nil {
		return nil, fmt.Errorf("parse %s: %w", filename, err)
	}

	return &tense, nil
}
