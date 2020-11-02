package mode

import (
"sync"
)

type mode struct {
	isInteractive bool
	isUuidMode bool
	isViewMode bool
}

func (m *mode) setInteractive() {
	m.isInteractive = true
}

func (m *mode) setNonInteractive() {
	m.isInteractive = false
}

func (m *mode) isInteractiveMode() bool {
	return m.isInteractive == true
}

func (m *mode) isNonInteractiveMode() bool {
	return m.isInteractive == false
}

func (m *mode) setNonUuidMode() {
	m.isUuidMode = false
}

func (m *mode) setNoneViewMode() {
	m.isViewMode = false
}

func (m *mode) setUuidMode() {
	m.isUuidMode = true
}

func (m *mode) setViewMode() {
	m.isViewMode = true
}

func (m *mode) getUuidMode() bool {
	return m.isUuidMode
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

func SetUuidMode() {
	GetMode().setUuidMode()
}

func SetNonUuidMode() {
	GetMode().setNonUuidMode()
}

func SetViewMode() {
	GetMode().setViewMode()
}

func SetNonViewMode() {
	GetMode().setNoneViewMode()
}

func GetUuidMode() bool {
	return GetMode().getUuidMode()
}