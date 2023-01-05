package test

import (
	"os"

	"github.com/qmsu/fyne/v2"
	"github.com/qmsu/fyne/v2/internal"
	"github.com/qmsu/fyne/v2/storage"
)

type testStorage struct {
	*internal.Docs
}

func (s *testStorage) RootURI() fyne.URI {
	return storage.NewURI("file://" + os.TempDir())
}

func (s *testStorage) docRootURI() (fyne.URI, error) {
	return storage.Child(s.RootURI(), "Documents")
}
