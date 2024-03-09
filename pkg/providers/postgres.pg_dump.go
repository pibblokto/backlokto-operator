package providers

import (
	"fmt"
	"os"

	pg "github.com/habx/pg-commands"
	"github.com/pibblokto/backlokto/pkg/types"
)

func PostgresPgDump(job *types.BackupJob) {

	var pgPass string

	if job.Spec.Password == "" {
		pgPass = os.Getenv("PG_PASS")
	} else {
		pgPass = job.Spec.Password
	}

	dump, err := pg.NewDump(&pg.Postgres{
		Host:     job.Spec.Host,
		Port:     job.Spec.Port,
		DB:       job.Spec.Database,
		Username: job.Spec.Username,
		Password: pgPass,
	})

	if err != nil {
		fmt.Println(err)
	}

	dump.SetFileName(fmt.Sprintf(`%v.sql`, dump.DB))
	dump.SetupFormat("p")

	dumpExec := dump.Exec(pg.ExecOptions{StreamPrint: true})
	if dumpExec.Error != nil {
		fmt.Println(dumpExec.Error.Err)
		fmt.Println(dumpExec.Output)
	} else {
		fmt.Printf("Dump was succesfull. Filename: %s\n", dumpExec.File)
		fmt.Println(dumpExec.File)
	}

	artifacts := types.NewArtifacts()
	artifacts.Filepath = dumpExec.File

	for _, target := range job.Targets {
		TargetsMap[target.Type](&target, artifacts)
	}

}
