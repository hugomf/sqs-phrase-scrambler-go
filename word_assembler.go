package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func receiveWords(wrapper *SQSWrapper) []string {

	//receive first message
	protoWord, err := wrapper.PopMessage()
	if err != nil {
		return []string{}
	}

	wordDefinition := strings.Split(protoWord, ":")
	if len(wordDefinition) != 4 {
		return nil
	}

	phraseID := wordDefinition[0]
	size, _ := strconv.Atoi(wordDefinition[1])
	pos, _ := strconv.Atoi(wordDefinition[2])
	word := wordDefinition[3]

	phrase := make([]string, size)
	phrase[pos] = word

	wg := &sync.WaitGroup{}

	for i := 1; i < size; i++ {
		wg.Add(1)
		go updateWord(wrapper, phraseID, phrase, wg)

	}
	wg.Wait()

	return phrase
}

func updateWord(wrapper *SQSWrapper, phraseID string, phrase []string, wg *sync.WaitGroup) {

	defer wg.Done()

	protoWord, err := wrapper.PopMessage()
	if err != nil {
		fmt.Println(err.Error())
	}

	wordDefinition := strings.Split(protoWord, ":")
	if len(wordDefinition) == 4 {
		if phraseID == wordDefinition[0] {
			pos, err := strconv.Atoi(wordDefinition[2])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return
			}
			phrase[pos] = wordDefinition[3]
		}
	}
}

func assemmbler_runner() {

	wrapper, err := NewSQSWrapper("phrase-scrambler-queue")
	if err != nil {
		fmt.Println(err.Error())
	}
	for {
		phrase := receiveWords(wrapper)
		if phrase != nil {
			fmt.Println("La frase final: " + strings.Join(phrase, ""))
		}
	}
}
