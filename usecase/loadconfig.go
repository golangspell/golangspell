package usecase

import "github.com/danilovalente/golangspell/domain"

//LoadConfig triggers the lazy loading of the application Config
func LoadConfig() {
	domain.GetConfig()
}
