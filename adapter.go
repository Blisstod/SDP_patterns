package main

import "fmt"

type PlateInKg struct {
	WeightInKg float64
}

func (p *PlateInKg) GetWeight() float64 {
	return p.WeightInKg
}

type PlateInLbs struct {
	WeightInLbs float64
}

func (p *PlateInLbs) GetWeight() float64 {
	return p.WeightInLbs
}

type KgToLbsAdapter struct {
	PlateInKg *PlateInKg
}

func NewKgToLbsAdapter(plateInKg *PlateInKg) *KgToLbsAdapter {
	return &KgToLbsAdapter{
		PlateInKg: plateInKg,
	}
}

func (a *KgToLbsAdapter) GetWeight() float64 {
	return a.PlateInKg.GetWeight() * 2.20462
}

func main() {
	plateInKg := &PlateInKg{20}
	adapter := NewKgToLbsAdapter(plateInKg)

	weightInLbs := adapter.GetWeight()
	fmt.Printf("Weight in pounds: %.2f lbs\n", weightInLbs)

}
