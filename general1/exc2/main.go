package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var hmArr = [7]string{
	"+---+\n" +
		"    |\n" +
		"    |\n" +
		"    |\n" +
		"  ===\n",
	"+---+\n" +
		" 0  |\n" +
		"    |\n" +
		"    |\n" +
		"  ===\n",
	"+---+\n" +
		" 0  |\n" +
		" |  |\n" +
		"    |\n" +
		"  ===\n",
	"+---+\n" +
		" 0  |\n" +
		"/|  |\n" +
		"    |\n" +
		"  ===\n",
	"+---+\n" +
		" 0  |\n" +
		"/|\\|\n" +
		"    |\n" +
		"  ===\n",
	"+---+\n" +
		" 0  |\n" +
		"/|\\|\n" +
		"/   |\n" +
		"  ===\n",
	"+---+\n" +
		" 0  |\n" +
		"/|\\|\n" +
		"/\\ |\n" +
		"  ===\n",
}

var wordArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER",
	"ZODIAC", "ZOMBY", "FLUFF",
}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func main() {

	fmt.Println(getRandomWord())
	//show Game board
	for {
		showBoard()
		guess := getUserLetter()

		if strings.Contains(randWord, guess) {

		}

	}

	// A. if they guessed letter in word
	// Add to correctLetters
	// 1. Are there more letters to guess
	// 2. If no more letters to guess (YOU WIN)
	// B. if they guessed letter not in word
	// 1. Add new letter to guessedLetters,
	// wrongGuesses
	// 2. Check if wrong

}

func getRandomWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)

	randWord = wordArr[rand.Intn(7)]

	correctLetters = make([]string, len(randWord))

	return randWord
}

func showBoard() {
	fmt.Println(hmArr[len(wrongGuesses)])
	fmt.Print("Secret Word : ")
	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}

	fmt.Print("\nIncorrect Guesses : ")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	fmt.Println("")

}

func getUserLetter() string {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nGuess a letter : ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)
		var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

		if len(guess) != 1 {
			fmt.Println("Please enter only one letter")
		} else if IsLetter(guess) {
			fmt.Println("Please enter a letter")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please enter a Letter you haven't guessed")
		} else {
			return guess
		}

	}

}
