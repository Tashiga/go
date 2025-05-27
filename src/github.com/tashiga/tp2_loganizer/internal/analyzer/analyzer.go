package analyzer

import (
	"errors"
	"math/rand"
	"os"
	"sync"
	"time"
)

type LogResult struct {
	LogID        string `json:"log_id"`
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

var ErrFileNotFound = errors.New("fichier introuvable")
var ErrParsing = errors.New("erreur de parsing")

func AnalyzeLog(logID, filePath string) LogResult {
	_, err := os.Stat(filePath)
	if err != nil {
		return LogResult{
			LogID:        logID,
			FilePath:     filePath,
			Status:       "FAILED",
			Message:      "Fichier introuvable.",
			ErrorDetails: err.Error(),
		}
	}

	rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(150)+50) * time.Millisecond)

	if rand.Intn(100) < 10 {
		return LogResult{
			LogID:        logID,
			FilePath:     filePath,
			Status:       "FAILED",
			Message:      "Erreur de parsing du log.",
			ErrorDetails: ErrParsing.Error(),
		}
	}

	return LogResult{
		LogID:    logID,
		FilePath: filePath,
		Status:   "OK",
		Message:  "Analyse terminée avec succès.",
	}
}

func AnalyzeLogsConcurrently(logs []struct {
	ID   string
	Path string
}) []LogResult {
	var wg sync.WaitGroup
	resultsChan := make(chan LogResult, len(logs))

	for _, log := range logs {
		wg.Add(1)

		go func(logID, path string) {
			defer wg.Done()
			res := AnalyzeLog(logID, path)
			resultsChan <- res
		}(log.ID, log.Path)
	}

	wg.Wait()
	close(resultsChan)

	results := make([]LogResult, 0, len(logs))
	for res := range resultsChan {
		results = append(results, res)
	}

	return results
}
