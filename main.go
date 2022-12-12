package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//--------------------------------------------------------------------------------------
	// FILE FINDER PROJECT
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("provide a directory")
		return
	}

	files, err := ioutil.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.Size() == 0 {
			name := file.Name()
			fmt.Println(name)
		}
	}
	//--------------------------------------------------------------------------------------
	//______________________________________________________________________________________
	// EXCHANGE RATIO
	// const (
	// 	ETH = 9 - iota
	// 	ADA
	// )
	// rates := [...]float64{
	// 	ETH :25.5,  // ethereum
	// 	ADA : 120.5, // cardano
	// }
	// fmt.Printf("1 BTC is %g ETH\n", rates[ETH])
	// fmt.Printf("1 BTC is %g ADA\n", rates[ADA])

	//______________________________________________________________________________________

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// MOODLY CHALLENGE
	// args := os.Args[1:]

	// if len(args) != 1 {
	// 	fmt.Println("[your name]")
	// 	return
	// }
	// name := args[0]

	// moods := [...]string{
	// 	"feels  sad",
	// 	"feels terrible",
	// 	"feels awesome",
	// 	"feels happy",
	// 	"feels good",
	// }

	// rand.Seed(time.Now().UnixNano())
	// n := rand.Intn(len(moods))

	// fmt.Printf("%s feels %s\n", name, moods[n])
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

	//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
	// #2 Project
	// const corpus = "" +
	// 	"Lady cat jumps again and again and again"

	// words := strings.Fields(corpus)
	// query := os.Args[1:]

	// queries:   // labeled break/continue for the parent loop
	// for _, q := range query {
	// 	search: // labeled loop for the nested loop
	// 	for i, w := range words {
	// 		switch q {
	// 		case "and", "or", "the":
	// 			break search
	// 		}
	// 		if q == w {
	// 			fmt.Printf("#%-2d: %q\n",
	// 		i+1, w)
	// 		continue queries
	// 		}
	// 	}
	// }
	//XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>
	// 1st Project

	// const (
	// 	maxTurns = 5
	// 	usage = `Welcome to the Luck Number Game!
	// 	The game will pick %d random numbers.
	// 	Your mission is to guess one of those numbers.
	// 	The greater your numbers are, the harder it gets.

	// 	Wanna play?
	// `

	// )

	// rand.Seed(time.Now().UnixNano())

	// args := os.Args[1:]
	// if len(args) != 1 {
	// 	fmt.Printf(usage, maxTurns)
	// 	return
	// }

	// guess, err := strconv.Atoi(args[0])
	// if err != nil {
	// 	fmt.Println("Not a number")
	// 	return
	// }

	// if guess
	// 	fmt.Println("Pick a positive number")
	// 	return
	// }

	// for turn := 0; turn < maxTurns; turn++
	// 	n := rand.Intn(guess + 1)

	// 	if n == guess {
	// 		fmt.Println("🏆 YOU WON!!!")
	// 		return
	// 	}
	// }
	// fmt.Println("🎴 YOU LOST... Try again")
	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>

}