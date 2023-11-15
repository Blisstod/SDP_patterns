package main

import "fmt"

type Lighting struct {
	isOn bool
}

func newLighting(isOn bool) *Lighting {
	return &Lighting{
		isOn: isOn,
	}
}

func (l *Lighting) TurnOn() {
	l.isOn = true
}

func (l *Lighting) TurnOff() {
	l.isOn = false
}

type Temperature struct {
	temp float64
}

func newTemperature(temp float64) *Temperature {
	return &Temperature{
		temp: temp,
	}
}

func (t *Temperature) setTemperature(temp float64) {
	t.temp = temp
	fmt.Printf("Temperature set to %f degrees.\n", temp)
}

type SecuritySystem struct {
}

func (ss *SecuritySystem) Activate() {
	fmt.Println("Security system activated.")
}

func (ss *SecuritySystem) Deactivate() {
	fmt.Println("Security system deactivated.")
}

type Home struct {
	kitchenLighting *Lighting
	hallLighting    *Lighting
	temperature     *Temperature
	ss              *SecuritySystem
}

func newHome() *Home {
	return &Home{
		kitchenLighting: newLighting(false),
		hallLighting:    newLighting(false),
		temperature:     newTemperature(25),
		ss:              &SecuritySystem{},
	}
}

func (h *Home) turnAllLightsOn() {
	h.hallLighting.TurnOn()
	h.kitchenLighting.TurnOn()
	fmt.Println("all lights in home have been turned on!")
}

func (h *Home) turnAllLightsOff() {
	h.hallLighting.TurnOff()
	h.kitchenLighting.TurnOff()
	fmt.Println("all lights in home have been turned off!")
}

func (h *Home) arriveHome() {
	h.turnAllLightsOn()
	h.temperature.setTemperature(24)
	h.ss.Deactivate()
}

func (h *Home) leaveHome() {
	h.turnAllLightsOff()
	h.temperature.setTemperature(21)
	h.ss.Activate()
}

func main() {
	home := newHome()
	home.leaveHome()
	home.arriveHome()
}
