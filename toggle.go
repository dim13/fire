package main

type Toggle struct {
	state bool
	On    func()
	Off   func()
}

func (t *Toggle) Toggle() {
	t.state = !t.state
	if t.state {
		t.On()
	} else {
		t.Off()
	}
}
