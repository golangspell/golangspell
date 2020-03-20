package config

import (
	"fmt"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	//configFileName defines the configuration file name
	configFileName = ".golangspell.json"
)

//Values stores the current configuration values
var (
	Values Config
	// Used for flags.
	CfgFile     string
	UserLicense string
	Author      string
	//DefautConfigFile holds the Golangspell's config file path
	DefautConfigFile string = ConfigFilePath
)

//Config contains the application's configuration values. Add here your own variables and bind it on init() function
type Config struct {
	//LogLevel - DEBUG or INFO or WARNING or ERROR or PANIC or FATAL
	LogLevel string
	//TestRun state if the current execution is a test execution
	TestRun bool
	//GoPath contains the configured or by convention GOPATH
	GoPath string
}

//GetHomeDir provides the User's Home Directory
func GetHomeDir() string {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	return home
}

//ConfigFilePath contains the path of the config file
var ConfigFilePath = fmt.Sprintf("%s%s%s", GetGolangspellHome(), PlatformSeparator, configFileName)

//GetGolangspellHome - platform agnostic
func GetGolangspellHome() string {
	home := GetHomeDir()
	return fmt.Sprintf("%s%s.golangspell", home, PlatformSeparator)
}

func init() {
	viper.BindEnv("TestRun", "TESTRUN")
	viper.SetDefault("TestRun", false)
	viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.SetDefault("LogLevel", "INFO")
	viper.BindEnv("GoPath", "GOPATH")
	viper.SetDefault("GoPath", fmt.Sprintf("%s%sgo", GetHomeDir(), PlatformSeparator))
	viper.Unmarshal(&Values)
}
