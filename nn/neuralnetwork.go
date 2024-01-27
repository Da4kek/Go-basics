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

func (nn *Neuralnetwork) Train(inputs []float64, targets []float64) {
}

func matMul(a, b [][]float64) [][]float64 {
	rowsA, colsA := len(a), len(a[0])
	rowsB, colsB := len(b), len(b[0])

	if colsA != rowsB {
		panic("Matrix dimensions do not match for multiplication")
	}

	result := make([][]float64, rowsA)
	for i := 0; i < rowsA; i++ {
		result[i] = make([]float64, colsB)
		for j := 0; j < colsB; j++ {
			sum := 0.0
			for k := 0; k < colsA; k++ {
				sum += a[i][k] * b[k][j]
			}
			result[i][j] = sum
		}
	}

	return result
}
