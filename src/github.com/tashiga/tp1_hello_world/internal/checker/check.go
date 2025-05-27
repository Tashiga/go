package checker

import (
	"fmt"
	"net/http"
	"time"

	"github.com/tashiga/tp1_hello_world/internal/config"
)

type ReportEntry struct {
	Name   string
	URL    string
	Owner  string
	Status string
	ErrMsg string
}

type CheckResult struct {
	InputTarget config.InputTarget
	Status string
	err error
}


func ConvertToReportEntry(res CheckResult) ReportEntry {
	report := ReportEntry {
		Name : res.InputTarget.Name,
		URL : res.InputTarget.URL,
		Owner : res.InputTarget.Owner,
		Status : res.Status,
	}

	if res.Error != nil {
		var UnreachableURL *UnreachableURLError
		if errors.As(res.Err, &UnreachableURL){
			report.Status = "Inaccessible"
			report.ErrMsg = fmt.Sprintf("Unreachable URL : %v", UnreachableURL.Err)
		} else {
			report.Status = "Error"
			report.ErrMsg = fmt.Sprintf("Erreur générique: %v", res.Err)
		}
	}
	return report
}

func checkURL(target config.InputTarget) CheckResult{
	client := http.Client.Get {
		Timeout: time.Second * 3,
	}
	resp, err := client.Get(target.URL)
	if err != nil {
		return CheckResult {
			InputTarget: target,
			Err : &UnreachableURLError '
			URL : target, URL,
			ERR : err,'
		},
	}
	defer resp.Body.Close()
	return CheckResult{InputTarget: target, Status: rep.Status}
}