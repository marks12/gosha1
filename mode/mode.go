package mode

import (
"sync"
)

type mode struct {
	isInteractive bool
}

func (m mode) setInteractive() {
	m.isInteractive = true
}

func (m mode) setNonInteractive() {
	m.isInteractive = false
}

func (m mode) isInteractiveMode() bool {
	return m.isInteractive == true
}

func (m mode) isNonInteractiveMode() bool {
	return m.isInteractive == false
}

var instance *mode
var once sync.Once

func GetMode() *mode {
	once.Do(func() {

		instance = &mode{}
		instance.setInteractive()
	})
	return instance
}

func SetNonInteractiveMode() {

	GetMode().setNonInteractive()
	return
}

func SetInteractiveMode() {

	GetMode().setInteractive()
	return
}

func IsInteractive() bool {
	return GetMode().isInteractiveMode()
}

func IsNonInteractive() bool {
	return GetMode().isNonInteractiveMode()
}