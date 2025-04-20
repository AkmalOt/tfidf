package service

import (
	"log"
	"math"
	"regexp"
	"strings"
)

func (s *ServiceImpl) Idf(txt string) map[string]float64 {
	sentenceRe := regexp.MustCompile(`[.!?]+`)
	sentences := sentenceRe.Split(txt, -1)

	var filteredSentences []string
	for _, sentence := range sentences {
		trimmed := strings.TrimSpace(sentence)
		if trimmed != "" {
			filteredSentences = append(filteredSentences, trimmed)
		}
	}

	log.Print(filteredSentences[0])
	wordRe := regexp.MustCompile(`[а-яёa-z0-9_]+`)
	allText := strings.ToLower(txt)
	words := wordRe.FindAllString(allText, -1)

	uniqueWords := make(map[string]bool)
	for _, word := range words {
		uniqueWords[word] = true
	}
	wordDocCount := make(map[string]float64)
	for word := range uniqueWords {
		var count float64
		for _, sentence := range filteredSentences {
			if strings.Contains(strings.ToLower(sentence), word) {
				count++
			}
		}
		wordDocCount[word] = count
	}
	totalDocs := float64(len(filteredSentences))
	idf := make(map[string]float64)
	for word, count := range wordDocCount {
		if count > 0 {
			idf[word] = math.Log(totalDocs / count)
		} else {
			idf[word] = 0
		}
	}
	return idf
}
