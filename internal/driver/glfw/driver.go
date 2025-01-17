// Package glfw provides a full Fyne desktop driver that uses the system OpenGL libraries.
// This supports Windows, Mac OS X and Linux using the gl and glfw packages from go-gl.
package glfw

import (
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/qmsu/fyne/v2"
	"github.com/qmsu/fyne/v2/internal/animation"
	intapp "github.com/qmsu/fyne/v2/internal/app"
	"github.com/qmsu/fyne/v2/internal/driver"
	"github.com/qmsu/fyne/v2/internal/driver/common"
	"github.com/qmsu/fyne/v2/internal/painter"
	intRepo "github.com/qmsu/fyne/v2/internal/repository"
	"github.com/qmsu/fyne/v2/storage/repository"
)

const mainGoroutineID = 1

var (
	curWindow *window
	isWayland = false
)

// Declare conformity with Driver
var _ fyne.Driver = (*gLDriver)(nil)

type gLDriver struct {
	windowLock sync.RWMutex
	windows    []fyne.Window
	device     *glDevice
	done       chan interface{}
	drawDone   chan interface{}

	animation *animation.Runner

	drawOnMainThread bool // A workaround on Apple M1, just use 1 thread until fixed upstream
}

func (d *gLDriver) RenderedTextSize(text string, textSize float32, style fyne.TextStyle) (size fyne.Size, baseline float32) {
	return painter.RenderedTextSize(text, textSize, style)
}

func (d *gLDriver) CanvasForObject(obj fyne.CanvasObject) fyne.Canvas {
	return common.CanvasForObject(obj)
}

func (d *gLDriver) AbsolutePositionForObject(co fyne.CanvasObject) fyne.Position {
	c := d.CanvasForObject(co)
	if c == nil {
		return fyne.NewPos(0, 0)
	}

	glc := c.(*glCanvas)
	return driver.AbsolutePositionForObject(co, glc.ObjectTrees())
}

func (d *gLDriver) Device() fyne.Device {
	if d.device == nil {
		d.device = &glDevice{}
	}

	return d.device
}

func (d *gLDriver) Quit() {
	if curWindow != nil {
		curWindow = nil
		fyne.CurrentApp().Lifecycle().(*intapp.Lifecycle).TriggerExitedForeground()
	}
	defer func() {
		recover() // we could be called twice - no safe way to check if d.done is closed
	}()
	close(d.done)
}

func (d *gLDriver) Run() {
	if goroutineID() != mainGoroutineID {
		panic("Run() or ShowAndRun() must be called from main goroutine")
	}
	d.runGL()
}

func (d *gLDriver) addWindow(w *window) {
	d.windowLock.Lock()
	defer d.windowLock.Unlock()
	d.windows = append(d.windows, w)
}

// a trivial implementation of "focus previous" - return to the most recently opened, or master if set.
// This may not do the right thing if your app has 3 or more windows open, but it was agreed this was not much
// of an issue, and the added complexity to track focus was not needed at this time.
func (d *gLDriver) focusPreviousWindow() {
	d.windowLock.RLock()
	wins := d.windows
	d.windowLock.RUnlock()

	var chosen fyne.Window
	for _, w := range wins {
		chosen = w
		if w.(*window).master {
			break
		}
	}

	if chosen == nil || chosen.(*window).view() == nil {
		return
	}
	chosen.RequestFocus()
}

func (d *gLDriver) windowList() []fyne.Window {
	d.windowLock.RLock()
	defer d.windowLock.RUnlock()
	return d.windows
}

func (d *gLDriver) initFailed(msg string, err error) {
	fyne.LogError(msg, err)

	if running() {
		d.Quit()
	} else {
		os.Exit(1)
	}
}

func goroutineID() int {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	// string format expects "goroutine X [running..."
	id := strings.Split(strings.TrimSpace(string(b)), " ")[1]

	num, _ := strconv.Atoi(id)
	return num
}

// NewGLDriver sets up a new Driver instance implemented using the GLFW Go library and OpenGL bindings.
func NewGLDriver() fyne.Driver {
	d := new(gLDriver)
	d.done = make(chan interface{})
	d.drawDone = make(chan interface{})
	d.animation = &animation.Runner{}

	repository.Register("file", intRepo.NewFileRepository())

	return d
}
