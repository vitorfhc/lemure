package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/vitorfhc/GoMario/gmcommon"
	"github.com/vitorfhc/GoMario/gmengine"
)

var engine gmengine.Engine
var config gmcommon.Configuration

// update is called every time in ebiten's
// game loop
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	engine.Update()

	ebitenutil.DebugPrint(screen, "Hello, World!")

	return nil
}

// handleFullscreen takes care of all ebiten configs
// in case fullscreen is on
func handleFullscreen() {
	if config.Fullscreen {
		config.Width, config.Height = ebiten.ScreenSizeInFullscreen()
		config.Scale = ebiten.DeviceScaleFactor()
		ebiten.SetFullscreen(true)
	}
}

func main() {
	// Set log flags to time only
	log.SetFlags(log.Ltime)

	// Gets all configurations and handles them
	config = gmcommon.LoadConfigurations()
	handleFullscreen()

	// Define start scene and creates Engine
	startScene := &gmengine.Scene{"Scene 01"}
	engine = gmengine.Engine{startScene}

	// Ebiten starts
	err := ebiten.Run(update, config.Width, config.Height, config.Scale, "GoMario")
	if err != nil {
		log.Fatal(err)
	}
}
