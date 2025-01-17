//go:build !windows && !android && !ios
// +build !windows,!android,!ios

package dialog

import (
	"path/filepath"

	"github.com/qmsu/fyne/v2"
	"github.com/qmsu/fyne/v2/storage"
	"github.com/qmsu/fyne/v2/theme"
)

func (f *fileDialog) getPlaces() []favoriteItem {
	lister, err := storage.ListerForURI(storage.NewURI("file:///"))
	if err != nil {
		fyne.LogError("could not create lister for /", err)
		return []favoriteItem{}
	}
	return []favoriteItem{favoriteItem{
		"Computer",
		theme.ComputerIcon(),
		lister,
	}}
}

func isHidden(file fyne.URI) bool {
	if file.Scheme() != "file" {
		fyne.LogError("Cannot check if non file is hidden", nil)
		return false
	}
	path := file.String()[len(file.Scheme())+3:]
	name := filepath.Base(path)
	return name == "" || name[0] == '.'
}

func hideFile(filename string) error {
	return nil
}

func fileOpenOSOverride(*FileDialog) bool {
	return false
}

func fileSaveOSOverride(*FileDialog) bool {
	return false
}
