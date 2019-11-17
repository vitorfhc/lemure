package lemure

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/vitorfhc/lemure/common"
)

// Engine is the type holding all Lemure Engine
// data and methods necessary to run
type Engine struct {
	CurrentScene *Scene                // current scene being displayed
	Config       *common.Configuration // the configurations for the engine
}

// Update is responsible for updating the current scene
func (e *Engine) update(screen *ebiten.Image) error {
	log.Println("Engine update")
	e.CurrentScene.Update()

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	black := color.RGBA{0, 0, 0, 0xff}
	screen.Fill(black)

	return nil
}

// Run is responsible for running the game loop
// properly after setting all configurations
func (e *Engine) Run() {
	log.Println("Run engine")

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
		log.Fatal(err)
	}
}

func (e *Engine) handleConfiguration() {
	e.Config.Scale = ebiten.DeviceScaleFactor()

	if e.Config.Fullscreen {
		e.Config.Width, e.Config.Height = ebiten.ScreenSizeInFullscreen()
		ebiten.SetFullscreen(true)
	}
}
