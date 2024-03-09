package types

type Specs struct {
	Port              int    `yaml:"port"`
	Path              string `yaml:"path"`
	Host              string `yaml:"host"`
	Database          string `yaml:"database"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	RdsRegion         string `yaml:"rds_region"`
	RdsIdentifier     string `yaml:"rds_identifier"`
	RdsSnapshotPrefix string `yaml:"rds_snapshot_prefix"`
	AccessKey         string `yaml:"access_key"`
	SecretKey         string `yaml:"secret_key"`
}

type Target struct {
	Type         string `yaml:"type"`
	S3BucketName string `yaml:"s3_bucket_name"`
	S3BucketKey  string `yaml:"s3_bucket_key"`
	Region       string `yaml:"region"`
	AccessKey    string `yaml:"access_key"`
	SecretKey    string `yaml:"secret_key"`
}

type BackupJob struct {
	Id       string            `yaml:"id"`
	Provider string            `yaml:"provider"`
	Spec     Specs             `yaml:"spec"`
	Targets  []Target          `yaml:"targets"`
	Tags     map[string]string `yaml:"tags"`
}

type Jobs struct {
	Jobs []BackupJob `yaml:"jobs"`
}

type Artifacts struct {
	Filepath string
}
