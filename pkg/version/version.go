package version

import (
	"strconv"
	"strings"
)

// version is set at build time (see builder/go-wrapper.sh)
var version string

func String() string {
	if version != "" {
		return version
	}
	return "dev"
}

// Release returns the release version (which is the version with the
// "-<commit>" suffix removed)
func Release() string {
	parts := strings.SplitN(String(), "-", 2)
	return parts[0]
}

func Dev() bool {
	return String() == "dev"
}

type Version struct {
	Dev       bool
	Date      string
	Iteration int
}

func (v *Version) Before(other *Version) bool {
	return v.Date < other.Date || v.Date == other.Date && v.Iteration < other.Iteration
}

func Parse(s string) *Version {
	if len(s) == 0 || s[0] != 'v' || len(s) < 11 {
		return &Version{Dev: true}
	}
	v := &Version{Date: s[1:9]}
	v.Iteration, _ = strconv.Atoi(s[10:])
	return v
}
