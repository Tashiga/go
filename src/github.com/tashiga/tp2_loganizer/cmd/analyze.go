package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tashiga/tp2_loganizer/internal/analyzer"
	"github.com/tashiga/tp2_loganizer/internal/config"
	"github.com/tashiga/tp2_loganizer/internal/reporter"
)

var configPath string
var outputPath string

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyse les fichiers de logs définis dans un fichier de config JSON",
	Run: func(cmd *cobra.Command, args []string) {
		configs, err := config.LoadConfig(configPath)
		if err != nil {
			fmt.Println("Erreur lors du chargement de la config :", err)
			return
		}
		fmt.Println("Lancement de l'analyse concurrente...")
		logsToAnalyze := make([]struct {
			ID   string
			Path string
		}, len(configs))
		for i, c := range configs {
			logsToAnalyze[i] = struct {
				ID   string
				Path string
			}{c.ID, c.Path}
		}
		results := analyzer.AnalyzeLogsConcurrently(logsToAnalyze)

		for _, r := range results {
			var pathErr *os.PathError
			if errors.As(errors.New(r.ErrorDetails), &pathErr) {
				fmt.Printf("ID: %s, Path: %s, Status: %s, Problème de chemin : %s\n",
					r.LogID, r.FilePath, r.Status, pathErr.Path)
			} else if errors.Is(errors.New(r.ErrorDetails), os.ErrNotExist) {
				fmt.Printf("ID: %s, Path: %s, Status: %s, Message: Fichier introuvable.\n",
					r.LogID, r.FilePath, r.Status)
			} else {
				fmt.Printf("ID: %s, Path: %s, Status: %s, Message: %s, ErrorDetails: %s\n",
					r.LogID, r.FilePath, r.Status, r.Message, r.ErrorDetails)
			}
		}

		if outputPath != "" {
			err := reporter.ExportJSON(results, outputPath)
			if err != nil {
				fmt.Println("Erreur lors de l'export JSON :", err)
			} else {
				fmt.Println("Rapport JSON exporté vers", outputPath)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
	analyzeCmd.Flags().StringVarP(&configPath, "config", "c", "", "Chemin du fichier de configuration JSON")
	analyzeCmd.Flags().StringVarP(&outputPath, "output", "o", "", "Chemin du fichier de sortie JSON")
	analyzeCmd.MarkFlagRequired("config")
}
