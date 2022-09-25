package cfg

import "fmt"

const (
	// VersionMajor major version
	VersionMajor = 0
	// VersionMinor minor version
	VersionMinor = 3
	// VersionPatch patch level
	VersionPatch = 1
)

type VersionInfo struct {
	Major     int
	Minor     int
	Patch     int
	BuildInfo string
}

func (v VersionInfo) String() string {
	return fmt.Sprintf("%v.%v.%v-%v", v.Major, v.Minor, v.Patch, BuildInfo)
}

func (v VersionInfo) IsNewer(o VersionInfo) bool {
	if o.Major < v.Major {
		return true
	}
	if o.Minor < v.Minor {
		return true
	}
	if o.Patch < v.Patch {
		return true
	}
	return false
}

var (
	// BuildInfo contains the build timestamp
	BuildInfo = "development"
	// Version info
	Version = VersionInfo{
		Major:     VersionMajor,
		Minor:     VersionMinor,
		Patch:     VersionPatch,
		BuildInfo: BuildInfo,
	}
)
