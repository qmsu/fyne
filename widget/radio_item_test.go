package widget

import (
	"testing"

	"github.com/qmsu/fyne/v2"
	"github.com/qmsu/fyne/v2/test"
	"github.com/qmsu/fyne/v2/theme"

	"github.com/stretchr/testify/assert"
)

func TestRadioItem_FocusIndicator_Centered_Vertically(t *testing.T) {
	item := newRadioItem("Hello", nil)
	render := test.WidgetRenderer(item).(*radioItemRenderer)
	render.Layout(fyne.NewSize(200, 100))

	focusIndicatorSize := theme.IconInlineSize() + theme.Padding()
	heightCenterOffset := (100 - focusIndicatorSize) / 2
	assert.Equal(t, fyne.NewPos(theme.Padding()/2, heightCenterOffset), render.focusIndicator.Position1)
}
