package network

import (
	"github.com/zond/neunet/neuron"
)

type Network struct {
	Inputs []*neuron.Neuron
	Outputs []*neuron.Neuron
	Neurons map[*neuron.Neuron]bool
}

func (self *Network) Signal(inputs []float64) (result []float64) {
	for index, _ := range inputs {		
		self.Inputs[index].Signal = &inputs[index]
	}
	for neuron, _ := range self.Neurons {
		neuron.Signal = nil
	}
	result = make([]float64, len(self.Outputs))
	for index, output := range self.Outputs {
		output.Signal = nil
		result[index] = output.Output()
	}
	return
}

func (self *Network) Teach(expected []float64) {
	for index, expect := range expected {
		self.Outputs[index].Teach(nil, expect)
	}
}