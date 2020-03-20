package config

import "runtime"

//PlatformIsWindows checks if the application is running on Windows
var PlatformIsWindows = getPlatformIsWindows()

// PlatformSeparator specific for the current platform
var PlatformSeparator = getPlatformSeparator()

func getPlatformIsWindows() bool {
	return runtime.GOOS == "windows"
}

func getPlatformSeparator() string {
	if PlatformIsWindows {
		return "\\"
	}
	return "/"

}
