package service

import (
	"regexp"
	"strings"
)

func (s *ServiceImpl) Tf(txt string) map[string]float64 {
	// ломается из-за доп знаков.
	//words := strings.Split(txt, documentType)

	// ломается из-за доп кириллицы.
	//re := regexp.MustCompile(`\b\w+\b`)
	re := regexp.MustCompile(`[а-яёa-z0-9_]+`)
	words := re.FindAllString(strings.ToLower(txt), -1)

	count := make(map[string]float64)
	for _, word := range words {
		count[word]++
	}

	totalWords := float64(len(words))
	for k, v := range count {
		count[k] = v / totalWords
	}
	return count
}
