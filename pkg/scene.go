package pkg

type IScene interface {
	Update(float32)
	Render()

	AddEntity(Drawer)
	RemoveEntity(Drawer) bool
}

type BaseScene struct {
	entities []Drawer
}

func MakeBaseScene() *BaseScene {
	return &BaseScene{}
}

func (s *BaseScene) Update(time_delta float32) {
}

func (s *BaseScene) Render() {
	for _, e := range s.entities {
		e.Draw()
	}
}

func (s *BaseScene) AddEntity(toAdd Drawer) {
	s.entities = append(s.entities, toAdd)
}

func (s *BaseScene) RemoveEntity(toRemove Drawer) bool {
	idx := -1
	for i, e := range s.entities {
		if e == toRemove {
			idx = i
			break
		}
	}

	if idx >= 0 {
		s.entities = append(s.entities[:idx], s.entities[idx+1:]...)
		return true
	}

	return false
}
