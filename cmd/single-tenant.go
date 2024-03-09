package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/pibblokto/backlokto/pkg/confparse"
	"github.com/pibblokto/backlokto/pkg/providers"
	"github.com/pibblokto/backlokto/pkg/types"
)

var ProvidersMap = map[string]func(*types.BackupJob){
	"postgres.pg_dump":         providers.PostgresPgDump,
	"aws.rds.snapshot":         providers.RdsSnapshot,
	"aws.rds.snapshot.cluster": providers.RdsClusterSnapshot,
	"mysql.mysqldump":          providers.MysqlDump,
}

func main() {

	configpathPtr := flag.String("config", "", "Path to config file with backup jobs declaration")
	tagsPtr := flag.String("tags", "", "List of tags to group jobs in the following format: KEY1=VALUE1,KEY2=VALUE2")
	flag.Parse()

	tagsMap := make(map[string]string)
	if *tagsPtr != "" {
		pairs := strings.Split(*tagsPtr, ",")

		// Iterate over key-value pairs
		for _, pair := range pairs {
			// Split pair by equal sign
			parts := strings.Split(pair, "=")
			if len(parts) == 2 {
				key := parts[0]
				value := parts[1]
				tagsMap[key] = value
			}
		}
	}

	if len(tagsMap) == 0 {
		var jobs types.Jobs = confparse.ParseConfig(*configpathPtr)

		// Running backup jobs
		for _, job := range jobs.Jobs {
			fmt.Printf("Running %s...", job.Id)
			ProvidersMap[job.Provider](&job)
		}
	} else {
		var allJobs types.Jobs = confparse.ParseConfig(*configpathPtr)
		var selectedJobs []types.BackupJob

		// Filtering jobs
		for _, job := range allJobs.Jobs {
			jobTags := job.Tags
			isMatch := true
			for key, value := range tagsMap {
				if jobTags[key] != value {
					isMatch = false
					break
				}
			}
			if isMatch {
				selectedJobs = append(selectedJobs, job)
			}
		}

		// Running filtered backup jobs
		for _, job := range selectedJobs {
			fmt.Printf("Running %s...", job.Id)
			ProvidersMap[job.Provider](&job)
		}
	}

}
