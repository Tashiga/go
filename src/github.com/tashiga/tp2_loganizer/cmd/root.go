package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "loganalyzer",
	Short: "LogAnalyzer est un outil CLI pour analyser des fichiers de logs.",
	Long: `LogAnalyzer permet d'analyser plusieurs fichiers de logs en parallèle,
de générer des rapports et de gérer les erreurs de manière robuste.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}
}
