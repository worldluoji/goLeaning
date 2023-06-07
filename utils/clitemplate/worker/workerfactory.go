package worker

import (
	"clitemplate/worker/gitworker"
	"clitemplate/worker/localcopyworker"

	"github.com/spf13/viper"
)

func GetWorker(workerType string) Worker {
	var worker Worker
	if workerType == "git" {
		worker = &gitworker.GitWorker{
			Url:    viper.Get("remoteAddr").(string),
			Branch: viper.Get("branch").(string),
		}
	} else if workerType == "localcopy" {
		worker = &localcopyworker.LocalCopyWorker{}
	}

	return worker
}
