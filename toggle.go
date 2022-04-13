package main

type Toggle struct {
	on  bool
	On  func()
	Off func()
}

func (t *Toggle) Toggle() {
	switch t.on = !t.on; t.on {
	case true:
		t.On()
	case false:
		t.Off()
	}
}
