package lemure

import "log"

type GameObject struct {
	Name string // the name of the GO
}

// Update is responsible for updating the state
// of the GO
// This should never be called by the user,
// it is exported so the Scene can use it
func (g *GameObject) Update() {
	log.Println("GO Update:", g.Name)
}
