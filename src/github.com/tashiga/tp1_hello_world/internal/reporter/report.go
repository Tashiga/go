package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tashiga/tp1_hello_world/internal/checker"
)

func ExportResultsToJsonFile(filePath string, results []checker.ReportEntry) error {
	data, err := json.MarshalIndent(results, "", " ")
	if err != nil {
		return fmt.Errorf("Impossible d'encoder les resultats en JSON : %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("Impossible d'écrire le rapport JSON dans le fichier %s : %w", filePath, err)
	}
	return nil
}
