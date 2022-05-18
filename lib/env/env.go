package main
//package env

import (
	"runtime"
)

func DetectRunningEnviroment() (ops string ,arc string, pnm string) {
	switch runtime.GOOS {
	case "windows":
		ops = "w"
	case "linux":
		ops = "l"
	case "darwin":
		ops = "m"
	}
	switch runtime.GOARCH {
	case "amd64":
		arc = "ia"
		pnm = "64"
	case "386":
		arc = "ia"
		pnm = "32"
	case "arm":
		arc = "am"
		pnm = "32"
	case "arm64":
		arc = "am"
		pnm = "64"
	}
	return ops,arc,pnm
}

func main() {
	DetectRunningEnviroment()
}