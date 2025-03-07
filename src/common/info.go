package common

import "runtime"

// build args
// go build -ldflags "\
//  -X github.com/makeopensource/leviathan/common.Version=0.1.0 \
//  -X github.com/makeopensource/leviathan/common.CommitInfo=$(git rev-parse HEAD) \
//  -X github.com/makeopensource/leviathan/common.BuildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ') \
//  -X github.com/makeopensource/leviathan/common.Platform=$(go env GOOS)/$(go env GOARCH) \
//  -X github.com/makeopensource/leviathan/common.Branch=$(git rev-parse --abbrev-ref HEAD) \
//  "
// main.go

var Version = "dev"
var CommitInfo = "dev"
var BuildDate = "dev"
var Branch = "dev" // Git branch
var GoVersion = runtime.Version()
