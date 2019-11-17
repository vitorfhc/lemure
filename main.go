package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/vitorfhc/lemure/common"
	"github.com/vitorfhc/lemure/engine"
)

var coreEngine engine.Engine
var config common.Configuration

// update is called every time in ebiten's
// game loop
func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Clears the screen
	black := color.RGBA{0, 0, 0, 0xff}
	screen.Fill(black)

	// Updates it
	coreEngine.Update()

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
	config = common.LoadConfigurations()
	handleFullscreen()

	// Define start scene and creates Engine
	startGameObjects := make([]engine.GameObject, 1, 1)
	startGameObjects[0] = engine.GameObject{Name: "Character"}
	startScene := &engine.Scene{Name: "Scene 01", GameObjects: startGameObjects}
	coreEngine = engine.Engine{CurrentScene: startScene}

	// Ebiten starts
	err := ebiten.Run(update, config.Width, config.Height, config.Scale, "lemure")
	if err != nil {
		log.Fatal(err)
	}
}
