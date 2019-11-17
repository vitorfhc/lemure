package gmengine

import "log"

type Scene struct {
	Name        string       // the name of the scene
	GameObjects []GameObject // list of gameobjects in the scene
}

// Update is responsible for updating every
// updatable object inside it
func (s *Scene) Update() {
	log.Println("Scene Update:", s.Name)

	for _, g := range s.GameObjects {
		g.Update()
	}
}
