package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	nwordsneeded := 50
	rand.Seed(time.Now().UnixNano())
	if !validateInput(os.Args) {
		return
	}
	trigrams := filereader(os.Args[1])
	start := initoutstr(trigrams)
	fmt.Println(generatetext(start, nwordsneeded, trigrams))

}

func validateInput(args []string) bool {
	if len(args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("trigrams <filepath>")
		return false
	}
	return true
}

func filereader(path string) []string {
	//Reads a file and returns a slice with all the trigrams in that file.
	//Open the input file for reading
	file, err := os.Open(path)
	check(err)
	defer file.Close()
	//store the string into Trigrams
	trigram := ""
	trigrams := make([]string, 0)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	//scan the first two words
	scanner.Scan()
	trigram = scanner.Text()
	scanner.Scan()
	trigram = trigram + " " + scanner.Text()
	//build a trigram for every word after the second
	for scanner.Scan() {
		trigram = trigram + " " + scanner.Text()
		trigrams = append(trigrams, trigram)
		trigram = trigram[strings.Index(trigram, " ")+1:]
	}
	check(scanner.Err())
	return trigrams
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func initoutstr(trigrams []string) string {
	//Randomly selects a trigram that starts with capital leter to start the generation
	startcanditates := make([]string, 0)
	r, _ := regexp.Compile("[A-Z]")
	for i := 0; i < len(trigrams); i++ {
		trigram := trigrams[i]
		if r.MatchString(trigram[:1]) {
			startcanditates = append(startcanditates, trigram)
		}
	}
	return PickRandomlyFromCanditates(startcanditates)
}

func PickRandomlyFromCanditates(canditates []string) string {
	if len(canditates) > 1 {
		randindx := rand.Intn(len(canditates))
		return canditates[randindx]
	} else if len(canditates) > 0 {
		return canditates[0]
	}
	return "error!"
}

func generatetext(start string, nwordsneeded int, trigrams []string) string {
	//Finds the matching trigrams and randomly selects one to continue, does this recursively until desired number of words is reached.
	nextcanditates, _ := getnextcanditates(start, trigrams)
	start = start + " " + PickRandomlyFromCanditates(nextcanditates)
	if len(strings.Fields(start)) < nwordsneeded {
		start = generatetext(start, nwordsneeded, trigrams)
	}
	return start
}

func getnextcanditates(start string, trigrams []string) ([]string, int) {
	words := strings.Fields(start)
	lastwords := words[len(words)-2] + " " + words[len(words)-1]
	nextcanditates := make([]string, 0)
	for _, trigram := range trigrams {
		//compare the last 2 words of the input text with the first two words of each digram and select canditates
		if lastwords == trigram[:strings.LastIndex(trigram, " ")] {
			nextword := trigram[strings.LastIndex(trigram, " ")+1:]
			nextcanditates = append(nextcanditates, nextword)
		}
	}
	return nextcanditates, len(nextcanditates)
}
