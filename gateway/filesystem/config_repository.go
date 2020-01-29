package filesystem

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"

	"github.com/danilovalente/golangspell/appcontext"
	"github.com/danilovalente/golangspell/config"
	"github.com/danilovalente/golangspell/domain"
)

func (repo ConfigRepository) ensureConfigPath() {
	golangspellhome := config.GetGolangspellHome()
	if _, err := os.Stat(golangspellhome); os.IsNotExist(err) {
		os.Mkdir(golangspellhome, 0700)
	}
}

func (repo ConfigRepository) ensureConfigFile() {
	repo.ensureConfigPath()
	if _, err := os.Stat(config.ConfigFilePath); err == nil {
		return
	}
	config := domain.BuildDefaultConfig()
	repo.Save(&config)
}

//ConfigRepository is the filesystem implementation of the repository for domain.ConfigRepository
type ConfigRepository struct {
}

//Get s the Config from the filesystem
func (repo ConfigRepository) Get() (*domain.Config, error) {
	repo.ensureConfigFile()
	configFile, err := os.Open(config.ConfigFilePath)
	if err != nil {
		panic(err)
	}
	defer configFile.Close()
	configData, err := ioutil.ReadAll(configFile)
	if err != nil {
		panic(err)
	}
	var config domain.Config
	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}

	return &config, nil
}

//Save s the Config in the filesystem
func (repo ConfigRepository) Save(configEntity *domain.Config) (string, error) {
	configData, err := json.MarshalIndent(configEntity, "", "  ")
	if err != nil {
		panic(err)
	}
	file, err := os.Create(config.ConfigFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = io.WriteString(file, string(configData))
	if err != nil {
		panic(err)
	}
	return config.ConfigFilePath, file.Sync()
}

//InitConfigRepository lazily loads a ConfigRepository
func InitConfigRepository() appcontext.Component {
	return ConfigRepository{}
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.ConfigRepository, InitConfigRepository)
}
