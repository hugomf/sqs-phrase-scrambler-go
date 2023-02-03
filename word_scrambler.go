package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
)

func scrambler_runner() {

	// wrapperReceive:= NewSQSWrapper("")

	inputPhrase := "Ya termine!!!"

	phraseID := uuid.New().String()
	words := []rune(inputPhrase)
	size := len(words)

	wordPositions := make([]string, size)
	for i, word := range words {
		wordPositions[i] = phraseID + ":" + strconv.Itoa(size) + ":" + strconv.Itoa(i) + ":" + string(word)
	}

	rand.Seed(time.Now().Unix())
	rand.Shuffle(size, func(i, j int) {
		wordPositions[i], wordPositions[j] = wordPositions[j], wordPositions[i]
	})

	wrapperSend, err := NewSQSWrapper("phrase-producer-queue")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	wg := &sync.WaitGroup{}

	for _, word := range wordPositions {
		wg.Add(1)
		go sendMessage(word, wrapperSend, wg)
	}
	wg.Wait()
}

func sendMessage(word string, wrapperSend *SQSWrapper, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println(word)
	_, err := wrapperSend.SendMessage(word)
	if err != nil {
		fmt.Println(err.Error())
	}

}
