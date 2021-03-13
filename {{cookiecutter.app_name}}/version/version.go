package version

import (
	"fmt"
	"runtime"
)

// GitCommit returns the git commit that was compiled. This will be filled in by the compiler.
var GitCommit string // nolint:gochecknoglobals

// Version returns the main version number that is being run at the moment.
const Version = "0.1.0"

// BuildDate returns the date the binary was built.
var BuildDate = "" // nolint:gochecknoglobals

// GoVersion returns the version of the go runtime used to compile the binary.
var GoVersion = runtime.Version() // nolint:gochecknoglobals

// OsArch returns the os and arch used to build the binary.
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH) // nolint:gochecknoglobals
