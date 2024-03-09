package providers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jamf/go-mysqldump"
	"github.com/pibblokto/backlokto/pkg/types"
)

func MysqlDump(job *types.BackupJob) {
	// Open connection to database
	var mysqlPass string

	if job.Spec.Password == "" {
		mysqlPass = os.Getenv("MYSQL_PASS")
	} else {
		mysqlPass = job.Spec.Password
	}

	config := mysql.NewConfig()
	config.User = job.Spec.Username
	config.Passwd = mysqlPass
	config.DBName = job.Spec.Database
	config.Net = "tcp"
	config.Addr = fmt.Sprintf("%s:%d", job.Spec.Host, job.Spec.Port) //"your-hostname:your-port"

	dumpDir := "./"                                        // you should create this directory
	dumpFilenameFormat := fmt.Sprintf("%v", config.DBName) // accepts time layout string and add .sql at the end of file

	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		fmt.Println("Error opening database: ", err)
		return
	}

	// Register database with mysqldump
	dumper, err := mysqldump.Register(db, dumpDir, dumpFilenameFormat)
	if err != nil {
		fmt.Println("Error registering databse:", err)
		return
	}

	fmt.Println()
	fmt.Println(dumpDir)
	fmt.Println(dumpFilenameFormat)
	fmt.Println()
	// Dump database to file
	err = dumper.Dump()
	if err != nil {
		fmt.Println("Error dumping:", err)
		return
	}
	fmt.Println(dumpFilenameFormat)
	fmt.Printf("File is saved to %s.sql\n", dumpFilenameFormat)

	// Close dumper, connected database and file stream.
	dumper.Close()

	artifacts := types.NewArtifacts()
	fmt.Printf("%s%s.sql\n", dumpDir, dumpFilenameFormat)
	artifacts.Filepath = fmt.Sprintf("%s%s.sql", dumpDir, dumpFilenameFormat)

	for _, target := range job.Targets {
		TargetsMap[target.Type](&target, artifacts)
	}
}
