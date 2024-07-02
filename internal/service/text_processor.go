package service

import (
	"regexp"
	"strings"
)

type Processor struct {
	groupMaxSentences int
}

func NewProcessor(groupMaxSentences int) *Processor {
	return &Processor{groupMaxSentences: groupMaxSentences}
}

func (s *Processor) SplitText(text string) []string {
	return s.splitTextIntoGroups(text)
}

func (s *Processor) splitTextIntoGroups(text string) []string {
	re := regexp.MustCompile(`[.!?]`)
	sentences := re.Split(text, -1)

	var groups []string
	var currentGroup []string

	for i := range sentences {
		sentences[i] = strings.TrimSpace(sentences[i])
		sentences[i] += ". "
	}

	sentences[len(sentences)-1] = strings.TrimRight(sentences[len(sentences)-1], " .")

	processedSentences := make([]string, 0, len(sentences))
	for i := range sentences {
		if sentences[i] == "" {
			continue
		}
		processedSentences = append(processedSentences, sentences[i])
	}

	for _, sentence := range processedSentences {
		currentGroup = append(currentGroup, strings.TrimSpace(sentence))

		if len(currentGroup) == s.groupMaxSentences {
			groups = append(groups, strings.Join(currentGroup, " "))
			currentGroup = nil
		}
	}

	if len(currentGroup) > 0 {
		groups = append(groups, strings.Join(currentGroup, " "))
	}

	return groups
}
