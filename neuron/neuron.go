package neuron

import (
	"math"
)

type Configuration interface {
	GetMomentum() float64
	GetRate() float64
	GetProportionalTerm() float64
}

type Axon struct {
	Src     *Neuron
	Dst     *Neuron
	Weights []float64
	Biases  []float64
	Outputs []float64
	DWeight float64
	DBias   float64
}

func (self *Axon) Output() float64 {
	return self.Source.Output()*self.Weight + self.Bias
}

func (self *Axon) teach(derr float64, age int) {
	dw := (derr /  
}

type Neuron struct {
	Config     Configuration
	Axons      []*Axon
	Generation uint64
	Signals    []float64
}

func (self *Neuron) Teach(generation uint64, age int, err float64) {
	derr := 0.0
	if generation > self.Generation && age < len(self.Signals) {
		self.Generation = generation
		oldInput := 0
		for _, axon := range self.Axons {
			oldInput += 
		}
		derr = (-self.Config.GetRate() / 2) + err*(1+self.Signals[len(self.Signals)-age])*(self.Signals[len(self.Sigmoids)-age])
		for _, axon := range self.Axons {
			axon.teach(age, derr)
		}
	}
}

func (self *Neuron) Output(generation uint64) float64 {
	if generation > self.Generation || len(self.Signals) == 0 {
		self.Generation = generation
		newSignal := 1.0
		for _, axon := range self.Axons {
			newSignal *= axon.Output()
		}
		self.Signals = append(self.Signals, (1-math.Pow(math.E, -newSignal))/(1+math.Pow(math.E, -newSignal)))
	}
	return self.Signals[len(self.Signals)-1]
}
