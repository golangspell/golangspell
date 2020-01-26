package usecase

import (
	"github.com/danilovalente/golangspell/appcontext"
)

//LoadConfig triggers the lazy loading of the application Config
func LoadConfig() {
	appcontext.Current.Get(appcontext.Config)
}
