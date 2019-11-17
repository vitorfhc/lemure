package lemure

import "log"

// Scene holds all gameobjects of that scene and
// updates them correctly
type Scene struct {
	Name        string        // the name of the scene
	GameObjects []*GameObject // list of gameobjects in the scene
}

// Update is responsible for updating every
// updatable object inside it
// This should never be called by the user,
// it is exported so the Engine can call it
func (s *Scene) Update() {
	log.Println("Scene Update:", s.Name)

	for _, g := range s.GameObjects {
		g.Update()
	}
}
