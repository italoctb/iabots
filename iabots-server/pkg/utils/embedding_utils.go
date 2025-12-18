package utils

import (
	"fmt"
	"math"
)

// CosineSimilarity calcula a similaridade entre dois vetores
func CosineSimilarity(a, b []float32) (float64, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("vector length mismatch")
	}
	var dot, normA, normB float64
	for i := range a {
		dot += float64(a[i]) * float64(b[i])
		normA += float64(a[i]) * float64(a[i])
		normB += float64(b[i]) * float64(b[i])
	}
	if normA == 0 || normB == 0 {
		return 0, nil
	}
	return dot / (math.Sqrt(normA) * math.Sqrt(normB)), nil
}

// NormalizeVector retorna o vetor normalizado
func NormalizeVector(v []float32) ([]float32, error) {
	var norm float64
	for _, val := range v {
		norm += float64(val * val)
	}
	if norm == 0 {
		return nil, fmt.Errorf("zero norm")
	}
	sqrtNorm := float32(math.Sqrt(norm))
	result := make([]float32, len(v))
	for i, val := range v {
		result[i] = val / sqrtNorm
	}
	return result, nil
}
