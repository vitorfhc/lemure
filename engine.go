package lemure

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/sirupsen/logrus"
	"github.com/vitorfhc/lemure/common"
	_ "github.com/vitorfhc/lemure/logger" // side effect is setting the logrus level
)

// Engine is the type holding all Lemure Engine
// data and methods necessary to run
type Engine struct {
	CurrentScene *Scene                // current scene being displayed
	Config       *common.Configuration // the configurations for the engine
}

// Update is responsible for updating the current scene
func (e *Engine) update(screen *ebiten.Image) error {
	logrus.Trace("update engine")
	e.CurrentScene.Update()

	if ebiten.IsDrawingSkipped() {
		logrus.WithField("skip", true).Trace("engine drawing")
		return nil
	}

	logrus.WithField("skip", false).Trace("engine drawing")
	black := color.RGBA{0, 0, 0, 0xff}
	screen.Fill(black)

	return nil
}

// Run is responsible for running the game loop
// properly after setting all configurations
func (e *Engine) Run() {
	logrus.Trace("engine run")

	e.Config = common.LoadConfigurations()
	e.handleConfiguration()

	err := ebiten.Run(
		e.update,
		e.Config.Width,
		e.Config.Height,
		e.Config.Scale,
		"Lemure Engine",
	)

	if err != nil {
		logrus.Fatal(err)
	}
}

func (e *Engine) handleConfiguration() {
	logrus.Trace("engine handling configuration")

	e.Config.Scale = ebiten.DeviceScaleFactor()
	logrus.WithField("device-scale-factor", e.Config.Scale).Debug("engine handling")

	if e.Config.Fullscreen {
		logrus.WithFields(logrus.Fields{
			"width":      e.Config.Width,
			"height":     e.Config.Height,
			"scale":      e.Config.Scale,
			"fullscreen": e.Config.Fullscreen,
		}).Debug("engine handling")
		e.Config.Width, e.Config.Height = ebiten.ScreenSizeInFullscreen()
		ebiten.SetFullscreen(true)
	}
}
