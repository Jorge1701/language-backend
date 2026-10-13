package generator

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

// TenseData contains a tense's display name and irregular conjugations.
type TenseData struct {
	Name           string                       `yaml:"name"`
	RegularEndings map[string]map[string]string `yaml:"regular_endings"`
	Irregular      map[string]map[string]string `yaml:"irregular"`
}

// LoadTenseData reads a YAML file in the name/irregular format.
func LoadTenseData(filename string) (*TenseData, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", filename, err)
	}

	var tense TenseData
	if err := yaml.Unmarshal(data, &tense); err != nil {
		return nil, fmt.Errorf("parse %s: %w", filename, err)
	}

	return &tense, nil
}
