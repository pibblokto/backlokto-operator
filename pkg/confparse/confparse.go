package confparse

import (
	"github.com/pibblokto/backlokto/pkg/types"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func ParseConfig(configPath string) types.Jobs {
	var jobs types.Jobs

	f, err := os.ReadFile(configPath)

	if err != nil {
		log.Fatalln(err)
	}

	if err := yaml.Unmarshal(f, &jobs); err != nil {
		log.Fatalln(err)
	}

	return jobs
}
