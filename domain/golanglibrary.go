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
	return fmt.Sprintf("%s%sbin%s%s", config.Values.GoPath, config.PlatformSeparator, config.PlatformSeparator, golangLibrary.Name)
}

//SrcPath of the library
func (golangLibrary *GolangLibrary) SrcPath() string {
	return fmt.Sprintf("%s%ssrc%s%s", config.Values.GoPath, config.PlatformSeparator, config.PlatformSeparator, golangLibrary.URLToPackage())
}

//URLToPackage returns the package name referenced by the URL
func (golangLibrary *GolangLibrary) URLToPackage() string {
	return strings.ReplaceAll(strings.ReplaceAll(golangLibrary.URL, "http://", ""), "https://", "")
}
