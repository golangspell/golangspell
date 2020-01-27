package domain

import (
	"fmt"
	"strings"

	"github.com/danilovalente/golangspell/config"
)

//GolangLibrary contains the deployment info of an installed Golang Library
type GolangLibrary struct {
	URL  string `json:"url"`
	Name string `json:"name"`
}

//BinPath of the library
func (golangLibrary *GolangLibrary) BinPath() string {
	return fmt.Sprintf("%s/bin/%s", config.Values.GoPath, golangLibrary.Name)
}

//SrcPath of the library
func (golangLibrary *GolangLibrary) SrcPath() string {
	return fmt.Sprintf("%s/src/%s", config.Values.GoPath, strings.ReplaceAll(strings.ReplaceAll(golangLibrary.URL, "http://", ""), "https://", ""))
}

//URLToPackage returns the package name referenced by the URL
func (golangLibrary *GolangLibrary) URLToPackage() string {
	return strings.ReplaceAll(strings.ReplaceAll(golangLibrary.URL, "http://", ""), "https://", "")
}
