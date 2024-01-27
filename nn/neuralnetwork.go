package nn

import (
	"math/rand"
)

type Neuralnetwork struct {
	inputNodes   int
	hiddenNodes  int
	outputNodes  int
	learningRate float64

	weightsInput  [][]float64
	weightsOutput [][]float64
}

func randomMatrix(rows, cols int) [][]float64 {
	matrix := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		matrix[i] = make([]float64, cols)
		for j := 0; j < cols; j++ {
			matrix[i][j] = rand.Float64()
		}
	}
	return matrix
}

func Network(inputNodes, hiddenNodes, outputNodes int, learningRate float64) *Neuralnetwork {
	nn := &Neuralnetwork{
		inputNodes:   inputNodes,
		hiddenNodes:  hiddenNodes,
		outputNodes:  outputNodes,
		learningRate: learningRate,
	}
	nn.weightsInput = randomMatrix(nn.hiddenNodes, nn.inputNodes)
	nn.weightsOutput = randomMatrix(nn.outputNodes, nn.hiddenNodes)
	return nn
}
