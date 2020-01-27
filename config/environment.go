package config

import (
	"fmt"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

//Values stores the current configuration values
var Values Config

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

func init() {
	viper.BindEnv("TestRun", "TESTRUN")
	viper.SetDefault("TestRun", false)
	viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.SetDefault("LogLevel", "INFO")
	viper.BindEnv("GoPath", "GOPATH")
	viper.SetDefault("GoPath", fmt.Sprintf("%s/go", GetHomeDir()))
	viper.Unmarshal(&Values)
}