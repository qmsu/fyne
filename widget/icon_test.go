package widget_test

import (
	"testing"

	"github.com/qmsu/fyne/v2"
	"github.com/qmsu/fyne/v2/layout"
	"github.com/qmsu/fyne/v2/test"
	"github.com/qmsu/fyne/v2/theme"
	"github.com/qmsu/fyne/v2/widget"
)

func TestIcon_Layout(t *testing.T) {
	test.NewApp()
	defer test.NewApp()

	for name, tt := range map[string]struct {
		resource fyne.Resource
	}{
		"empty": {},
		"resource": {
			resource: theme.CancelIcon(),
		},
	} {
		t.Run(name, func(t *testing.T) {
			icon := &widget.Icon{
				Resource: tt.resource,
			}

			window := test.NewWindow(fyne.NewContainerWithLayout(layout.NewCenterLayout(), icon))
			window.Resize(icon.MinSize().Max(fyne.NewSize(150, 200)))

			test.AssertRendersToMarkup(t, "icon/layout_"+name+".xml", window.Canvas())

			window.Close()
		})
	}
}
