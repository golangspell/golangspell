package usecase

import "github.com/golangspell/golangspell/domain"

//LoadConfig triggers the lazy loading of the application Config
func LoadConfig() {
	domain.GetConfig()
}
