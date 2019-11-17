package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/vitorfhc/GoMario/gmcommon"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	ebitenutil.DebugPrint(screen, "Hello, World!")

	return nil
}

func main() {
	config := gmcommon.LoadConfigurations()

	if config.Fullscreen {
		config.Width, config.Height = ebiten.ScreenSizeInFullscreen()
		config.Scale = ebiten.DeviceScaleFactor()
		ebiten.SetFullscreen(true)
	}

	err := ebiten.Run(update, config.Width, config.Height, config.Scale, "GoMario")
	if err != nil {
		log.Fatal(err)
	}
}
