//go:build !ios
// +build !ios

package dialog

import (
	"os"

	"github.com/qmsu/fyne/v2"
	"github.com/qmsu/fyne/v2/storage"
)

func getFavoriteLocations() (map[string]fyne.ListableURI, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	homeURI := storage.NewFileURI(homeDir)

	favoriteNames := append(getFavoriteOrder(), FavoriteOrder{Name: "Home"})
	favoriteLocations := make(map[string]fyne.ListableURI)
	for _, favName := range favoriteNames {
		var uri fyne.URI
		var err1 error
		if favName.Name == "Home" {
			uri = homeURI
		} else {
			uri, err1 = storage.Child(homeURI, favName.Name)
		}
		if err1 != nil {
			err = err1
			continue
		}

		listURI, err1 := storage.ListerForURI(uri)
		if err1 != nil {
			err = err1
			continue
		}
		favoriteLocations[favName.Name] = listURI
	}

	return favoriteLocations, err
}
