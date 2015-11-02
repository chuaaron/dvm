// +build windows

package main

import (
	"strings"
)

const dockerOS string = "Windows"
const binaryFileExt string = ".exe"

func getCleanDvmPathRegex() string {
	versionDir := getVersionDir("")
	escapedVersionDir := strings.Replace(versionDir, `\`, `\\`, -1)
	return escapedVersionDir + `\\(\d+\.\d+\.\d+|experimental);`
}

func validateShellFlag() {
	if shell != "powershell" && shell != "cmd" {
		die("The --shell flag or SHELL environment variable must be set when running on Windows. Available values are powershell and cmd.", nil, retCodeInvalidArgument)
	}
}