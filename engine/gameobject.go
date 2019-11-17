package engine

import "log"

type GameObject struct {
	Name string // the name of the GO
}

// Update is responsible for updating the state
// of the GO
func (g *GameObject) Update() {
	log.Println("GO Update:", g.Name)
}
