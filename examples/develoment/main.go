package main

import (
	"github.com/vitorfhc/lemure"
)

func main() {
	// creates gameobjects
	gameObject1 := &lemure.GameObject{
		Name: "GameObject 01",
	}
	gameObject2 := &lemure.GameObject{
		Name: "GameObject 02",
	}

	// creates a scene and inserts the gameobject in it
	scene := &lemure.Scene{
		Name: "New Scene",
	}
	scene.GameObjects = append(scene.GameObjects, gameObject1)
	scene.GameObjects = append(scene.GameObjects, gameObject2)

	// creates the engine and set the current scene
	engine := &lemure.Engine{}
	engine.CurrentScene = scene

	// runs the game loop
	engine.Run()
}
