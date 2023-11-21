package main

import "fmt"

type Lighting struct {
	isOn bool
}

func (l *Lighting) turnOn() {
	l.isOn = true
}

func (l *Lighting) turnOff() {
	l.isOn = false
}

type Temperature struct {
	temp int
}

func (t *Temperature) setTemperature(temp int) {
	t.temp = temp
}

type SecuritySystem struct{}

func (ss *SecuritySystem) activate() {
	fmt.Println("Security System has been activated!")
}

func (ss *SecuritySystem) deactivate() {
	fmt.Println("Security System has been deactivated!")
}

type Home struct {
	kitchenLighting *Lighting
	hallLighting    *Lighting
	temperature     *Temperature
	ss              *SecuritySystem
}

func newHome() *Home {
	return &Home{
		kitchenLighting: &Lighting{isOn: false},
		hallLighting:    &Lighting{isOn: false},
		temperature:     &Temperature{temp: 24},
		ss:              &SecuritySystem{},
	}
}

func (h *Home) arriveHome() {
	h.ss.deactivate()
	h.temperature.setTemperature(24)
	h.kitchenLighting.turnOn()
	h.hallLighting.turnOn()
}

func (h *Home) leaveHome() {
	h.ss.activate()
	h.temperature.setTemperature(21)
	h.kitchenLighting.turnOff()
	h.hallLighting.turnOff()
}

func main() {
	home := newHome()
	home.arriveHome()
	home.leaveHome()
}
