package main

import (
	"fmt"
	"strconv"
	"strings"
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

	for i := 1; i < size; i++ {

		protoWord, err := wrapper.PopMessage()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		wordDefinition := strings.Split(protoWord, ":")
		if len(wordDefinition) != 4 {
			return nil
		}
		if phraseID == wordDefinition[0] {
			pos, err := strconv.Atoi(wordDefinition[2])
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				return nil
			}
			phrase[pos] = wordDefinition[3]
		}

	}

	return phrase
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
