package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func scrambler_runner() {

	// wrapperReceive:= NewSQSWrapper("")

	inputPhrase := "Hello World!"

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

	wrapperSend, err := NewSQSWrapper("phrase-scrambler-queue")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, word := range wordPositions {
		fmt.Println(word)
		_, err := wrapperSend.SendMessage(word)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		// fmt.Println(*resp.MessageId)
	}
}
