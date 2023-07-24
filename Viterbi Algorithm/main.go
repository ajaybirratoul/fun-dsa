package main

import (
	"fmt"
	"math"
)

type HMM struct {
	States      []string
	Observation []string
	StartProb   map[string]float64
	TransProb   map[string]map[string]float64
	EmitProb    map[string]map[string]float64
}

func (hmm *HMM) Viterbi(obs []string) []string {
	n := len(obs)
	T := len(hmm.States)

	dp := make([][]float64, n)
	path := make([][]int, n)

	for i := range dp {
		dp[i] = make([]float64, T)
		path[i] = make([]int, T)
	}

	// Initialize the first column (t = 0)
	for j, state := range hmm.States {
		dp[0][j] = hmm.StartProb[state] * hmm.EmitProb[state][obs[0]]
		path[0][j] = -1
	}

	// Compute the highest probability for each state and store the path
	for t := 1; t < n; t++ {
		for j, state := range hmm.States {
			maxProb := math.Inf(-1)
			maxIndex := -1

			for i, prevState := range hmm.States {
				prob := dp[t-1][i] * hmm.TransProb[prevState][state] * hmm.EmitProb[state][obs[t]]
				if prob > maxProb {
					maxProb = prob
					maxIndex = i
				}
			}

			dp[t][j] = maxProb
			path[t][j] = maxIndex
		}
	}

	// Find the best final state (maximum probability in the last column)
	maxProb := math.Inf(-1)
	maxIndex := -1
	for i := range hmm.States {
		prob := dp[n-1][i]
		if prob > maxProb {
			maxProb = prob
			maxIndex = i
		}
	}

	// Trace back the path to find the most likely sequence of hidden states
	result := make([]string, n)
	result[n-1] = hmm.States[maxIndex]
	for t := n - 2; t >= 0; t-- {
		maxIndex = path[t+1][maxIndex]
		result[t] = hmm.States[maxIndex]
	}

	return result
}

func main() {
	hmm := HMM{
		States:      []string{"Healthy", "Fever"},
		Observation: []string{"normal", "cold", "dizzy"},
		StartProb: map[string]float64{
			"Healthy": 0.6,
			"Fever":   0.4,
		},
		TransProb: map[string]map[string]float64{
			"Healthy": {
				"Healthy": 0.7,
				"Fever":   0.3,
			},
			"Fever": {
				"Healthy": 0.4,
				"Fever":   0.6,
			},
		},
		EmitProb: map[string]map[string]float64{
			"Healthy": {
				"normal": 0.5,
				"cold":   0.4,
				"dizzy":  0.1,
			},
			"Fever": {
				"normal": 0.1,
				"cold":   0.3,
				"dizzy":  0.6,
			},
		},
	}

	observation := []string{"normal", "cold", "dizzy"}
	result := hmm.Viterbi(observation)

	fmt.Println("Most likely sequence of hidden states:", result)
}
