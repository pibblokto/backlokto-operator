package providers

import (
	"github.com/pibblokto/backlokto/pkg/targets"
	"github.com/pibblokto/backlokto/pkg/types"
)

var TargetsMap = map[string]func(*types.Target, *types.Artifacts){
	"s3": targets.S3Target,
}
