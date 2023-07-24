# Viterbi Algorithm
This is a Go implementation of the Viterbi Algorithm, a dynamic programming algorithm used to find the most likely sequence of hidden states in a Hidden Markov Model (HMM). The Viterbi Algorithm is widely used in various applications, including speech recognition, natural language processing, and bioinformatics.

## How the Viterbi Algorithm Works

The Viterbi Algorithm operates on a Hidden Markov Model (HMM), which consists of a set of hidden states, observable states (observations), and probabilities for transitioning between states and emitting observations. The algorithm uses dynamic programming to efficiently find the most probable sequence of hidden states given a sequence of observations.

## Usage

To use the Viterbi Algorithm, follow these steps:

1. Create a new Go project or add the `main.go` file containing the Viterbi Algorithm implementation to your existing project.

2. Define your Hidden Markov Model (HMM) with the states, observations, start probabilities, transition probabilities, and emission probabilities.

3. Create an instance of the `HMM` struct and call the `Viterbi` method, passing the sequence of observations to find the most likely sequence of hidden states.

Here's an example of how to use the Viterbi Algorithm:

```go
func main() {
    hmm := HMM{
        States:      []string{"Healthy", "Fever"},
        Observations: []string{"normal", "cold", "dizzy"},
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

    observations := []string{"normal", "cold", "dizzy"}
    result := hmm.Viterbi(observations)

    fmt.Println("Most likely sequence of hidden states:", result)
}
```

## Features

- Hidden Markov Model (HMM): The implementation includes a simple representation of the Hidden Markov Model with states, observations, and probabilities for efficient processing.

- Dynamic Programming: The Viterbi Algorithm uses dynamic programming to efficiently find the most likely sequence of hidden states.
