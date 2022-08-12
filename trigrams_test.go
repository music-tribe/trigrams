package main

import (
	"math"
	"os"
	"strings"
	"testing"
)

func TestTheNumberOfTrigramsIsNMinusTwoForTextWithNWords(t *testing.T) {
	trigrams := filereader("sample.txt")
	fileContents, _ := os.ReadFile("sample.txt")
	numOfTrigrams := len(trigrams)
	numOfWordsInText := len(strings.Fields(string(fileContents)))
	if numOfTrigrams != numOfWordsInText-2 {
		t.Errorf("There are %d trigrams when %d was expected based on the number of words which is %d", numOfTrigrams, numOfWordsInText-2, numOfWordsInText)
	}
}
func TestAllTheTrigramsExistAtLeastOnceInTheOriginalText(t *testing.T) {
	trigrams := filereader("sample.txt")
	fileContents, _ := os.ReadFile("sample.txt")
	fileContentsOneLine := strings.ReplaceAll(string(fileContents), "\r\n", " ")
	fileContentsOneLine = removeDoubleSpaces(fileContentsOneLine)

	for _, trigram := range trigrams {
		if !strings.Contains(fileContentsOneLine, trigram) {
			t.Errorf("Trigram does not exist in the original text: %s", trigram)
		}
	}
}

func TestFrequencyOfWordsIsPreserved(t *testing.T) {
	//confirms the preservense of frequency within 2%
	trigrams := filereader("sample.txt")
	for _, trigram := range trigrams {
		nextcanditates, numOfCanditates := getnextcanditates(trigram, trigrams)
		if numOfCanditates > 1 {
			canditatesFrequencies := make(map[string]int)
			pickerFrequencies := make(map[string]int)
			for _, canditate := range nextcanditates {
				canditatesFrequencies[canditate]++
			}
			for i := 0; i < numOfCanditates*10000; i++ {
				pickerFrequencies[PickRandomlyFromCanditates(nextcanditates)]++
			}
			for word, freq := range canditatesFrequencies {
				candWordRatio := float64(freq) / float64(numOfCanditates) * 100
				actualWordRatio := float64(pickerFrequencies[word]) / float64(numOfCanditates*10000) * 100
				if math.Abs((actualWordRatio - candWordRatio)) > 2 {
					t.Errorf("Canditate word ration = %f actual word ratio = %f for word %s \n", candWordRatio, actualWordRatio, word)
				}

			}

		}

	}
}

func TestItGeneratesNwordsEachTime(t *testing.T) {
	N := 100
	trigrams := filereader("sample.txt")
	for i := 0; i < 1000; i++ {
		start := initoutstr(trigrams)
		generatedText := generatetext(start, N, trigrams)
		if strings.Contains(generatedText, "error!") {
			t.Errorf("Hit a dead end with text: %s", generatedText)
		}
		wc := len(strings.Fields(generatedText))
		if wc != N {
			t.Errorf("Generated %d words when expected %d words, the generated text was %s", wc, N, generatedText)
		}
	}
}

func removeDoubleSpaces(inputText string) string {
	tempText := inputText
	tempText = strings.ReplaceAll(tempText, "  ", " ")
	if len(tempText) != len(inputText) {
		tempText = removeDoubleSpaces(tempText)
	}
	return tempText
}
