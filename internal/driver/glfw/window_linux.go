package glfw

import "github.com/qmsu/fyne/v2"

func (w *window) platformResize(canvasSize fyne.Size) {
	w.canvas.Resize(canvasSize)
	w.canvas.Refresh(w.canvas.content)
}
