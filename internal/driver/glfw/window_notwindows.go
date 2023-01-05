//go:build !windows
// +build !windows

package glfw

import (
	"github.com/qmsu/fyne/v2"
	"github.com/qmsu/fyne/v2/internal"
)

func (w *window) setDarkMode() {
}

func (w *window) computeCanvasSize(width, height int) fyne.Size {
	return fyne.NewSize(internal.UnscaleInt(w.canvas, width), internal.UnscaleInt(w.canvas, height))
}
