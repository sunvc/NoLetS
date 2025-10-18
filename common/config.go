package common

import (
	_ "embed"
)

var LocalConfig = &Config{}

func SetDefaultVersionOrCommID(version, buildDate, commID string) {
	if len(version) > 0 {
		LocalConfig.System.Version = version
	} else {
		LocalConfig.System.Version = "1.0.0"
	}
	if len(commID) > 0 {
		LocalConfig.System.CommitID = commID
	} else {
		LocalConfig.System.CommitID = "f7efb70"
	}
	if len(buildDate) > 0 {
		LocalConfig.System.BuildDate = buildDate
	} else {
		LocalConfig.System.BuildDate = "2025-01-01 09:20:33"
	}
}
