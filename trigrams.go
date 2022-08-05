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

func validateInput(argcount int) bool {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("trigrams <filepath>")
		return false
	}
	return true
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func generatetext (start string, n int, trigrams []string) string{
	//Find the matching trigrams and randomly select one to continue, do this recursively until desired Number of words is reached.
	var output string
	nextcanditates:=make([]string,0)
	words := strings.Fields(start)
	lastwords := words[len(words)-2] + " " + words[len(words)-1]
	for _, trigram := range trigrams {
		if lastwords == trigram[:strings.LastIndex(trigram, " ")] {
			nextcanditates = append(nextcanditates, trigram[strings.LastIndex(trigram, " ")+1:])
		}
	}
	if len(nextcanditates)>1 {
		randindx := rand.Intn(len(nextcanditates)-1)
		output = start + " " + nextcanditates[randindx]
	} else {
		output = start + " " + nextcanditates[0]
	}
	if len(strings.Fields(output)) < n {
		output = generatetext(output,n,trigrams)
	}
	return output

}

func main() {
	//Validated input
	if validateInput(len(os.Args)) {
		//Open the input file for reading
		file, err := os.Open(os.Args[1])
		check(err)
		defer file.Close()
		//store the string into Trigrams
		trigram := ""
		trigrams := make([]string, 0)
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		//scan the first two words
		scanner.Scan()
		trigram =scanner.Text()
		scanner.Scan()
		trigram = trigram + " "+ scanner.Text()
		//build a trigram for every word after the second
		for scanner.Scan() {
			trigram = trigram + " " + scanner.Text()
			trigrams = append(trigrams, trigram)
			trigram = trigram[strings.Index(trigram, " ")+1:]
			//fmt.Println(scanner.Text())
		}
		check(scanner.Err())
		//Randomly select a trigram that starts with capital leter to start the generation
		var start string
		startcanditates:=make([]string,0)
		r, _ := regexp.Compile("[A-Z]")
		for i :=0; i<len(trigrams); i++ {
			trigram := trigrams[i]
			if r.MatchString(trigram[:1]){
				startcanditates = append(startcanditates, trigram)
			}
		}
		rand.Seed(time.Now().UnixNano())
		if len(startcanditates)>1{
			randindx := rand.Intn(len(startcanditates)-1)
			start = startcanditates[randindx]
		} else {
			start = startcanditates[0]
		}
		
		//Continue to find matching strings
		n:=100
		fmt.Println(generatetext(start, n,trigrams))

	}

}
