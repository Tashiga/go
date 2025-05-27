package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type InputTarget struct {
	Name  string
	URL   string
	Owner string
}

func LoadTargetFromFile(filePath string) ([]InputTarget, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("impossible de lire le fichier %s : %n", filePath, err)
	}
	var targets []InputTarget
	if err := json.Unmarshal(data, &targets); err != nil {
		return nil, fmt.Errorf("Impossible de lire le fichier %s : %n", filePath, err)
	}
	return targets, nil
}

func SaveTargetsToFile(filePath string, targets []InputTarget) error {
	data, err := json.MarshalIndent(targets, "", " ")
	if err != nil {
		return fmt.Errorf("impossible de sérialiser les données %s : %n", filePath, err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("impossible de créer/lire le fichier %s : %n", filePath, err)
	}
	return nil
}
