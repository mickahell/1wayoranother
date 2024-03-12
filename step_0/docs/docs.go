package docs

import (
	_ "embed"
)

//go:embed VERSION.txt
var version_file string

const app_name string = "Step 0"

func GetVersion() string {
	return version_file
}

func GetAppName() string {
	return app_name
}
