package main

import (
	"github.com/unixpickle/leea"
	"github.com/unixpickle/leea/demos/lightrnn"
	"github.com/unixpickle/sgd"
	"github.com/unixpickle/weakai/rnn/seqtoseq"
)

type Evaluator struct{}

func (_ Evaluator) Evaluate(e leea.Entity, s sgd.SampleSet) float64 {
	var seqSamples sgd.SliceSampleSet
	var totalLen int
	for i := 0; i < s.Len(); i++ {
		sample := s.GetSample(i).(Sample)
		seqSamples = append(seqSamples, seqtoseq.Sample{
			Inputs:  sample.InSeq(),
			Outputs: sample.OutSeq(),
		})
		totalLen += len(sample.InSeq())
	}
	r := e.(*Entity).RNN
	return -float64(lightrnn.Cost(r, seqSamples).(float32))
}
