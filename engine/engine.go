package engine

import "log"

type Engine struct {
	CurrentScene *Scene // current scene being displayed
}

// Update is responsible for updating the current scene
func (e *Engine) Update() {
	log.Println("Engine update")
	e.CurrentScene.Update()
}
