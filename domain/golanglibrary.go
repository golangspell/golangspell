package domain

import (
	"fmt"
	"strings"

	"github.com/golangspell/golangspell/config"
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

//ModPath of the library
func (golangLibrary *GolangLibrary) InstalledVersion() (string, error) {
	configRepo := GetConfigRepository()
	currentConfig, err := configRepo.Get()
	if err != nil {
		return "", fmt.Errorf("an error occurred while trying to read the config file")
	}
	if librarySpell, ok := currentConfig.Spellbook[golangLibrary.Name]; ok {
		return librarySpell.Version, nil
	} else {
		return "", fmt.Errorf("Spell %s not found", golangLibrary.Name)
	}
}

//ModPath of the library
func (golangLibrary *GolangLibrary) ModPath() string {
	if version, err := golangLibrary.InstalledVersion(); err == nil {
		return fmt.Sprintf("%s%spkg%smod%s%s", config.Values.GoPath, config.PlatformSeparator, config.PlatformSeparator, config.PlatformSeparator, golangLibrary.URLToPackage()+"@v"+version)
	} else {
		fmt.Println(err.Error())
		return ""
	}
}

//CachePath of the library
func (golangLibrary *GolangLibrary) CachePath() string {
	return fmt.Sprintf("%s%spkg%smod%scache%sdownload%s%s", config.Values.GoPath, config.PlatformSeparator, config.PlatformSeparator, config.PlatformSeparator, config.PlatformSeparator, config.PlatformSeparator, golangLibrary.URLToPackage())
}

//URLToPackage returns the package name referenced by the URL
func (golangLibrary *GolangLibrary) URLToPackage() string {
	return strings.ReplaceAll(strings.ReplaceAll(golangLibrary.URL, "http://", ""), "https://", "")
}
